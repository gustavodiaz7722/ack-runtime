package runtime

import (
	"fmt"
	"strings"

	ackcfg "github.com/aws-controllers-k8s/runtime/pkg/config"
	acktypes "github.com/aws-controllers-k8s/runtime/pkg/types"
	"github.com/awslabs/operatorpkg/status"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrlrt "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
)

var currentStatusGVK schema.GroupVersionKind

type fixedKindAccessor struct{}

func (*fixedKindAccessor) SetGroupVersionKind(schema.GroupVersionKind) {}
func (*fixedKindAccessor) GroupVersionKind() schema.GroupVersionKind   { return currentStatusGVK }

type gvkBoundUnstructured struct{ unstructured.Unstructured }

func (*gvkBoundUnstructured) GetObjectKind() schema.ObjectKind { return &fixedKindAccessor{} }

func NewStatusReconciler(
	cfg ackcfg.Config,
	rmf acktypes.AWSResourceManagerFactory,
	rd acktypes.AWSResourceDescriptor,
) acktypes.StatusReconciler {
	return &statusReconciler{
		cfg: cfg,
		rmf: rmf,
		rd:  rd,
	}
}

type statusReconciler struct {
	cfg ackcfg.Config
	rmf acktypes.AWSResourceManagerFactory
	rd  acktypes.AWSResourceDescriptor
}

func (r *statusReconciler) BindControllerManager(mgr ctrlrt.Manager) error {
	kind := r.rd.GroupVersionKind().Kind
	gvk := r.rd.GroupVersionKind()
	currentStatusGVK = gvk

	eventRec := mgr.GetEventRecorderFor(fmt.Sprintf("operatorpkg.%s.status", strings.ToLower(kind)))

	sc := status.NewGenericObjectController[*gvkBoundUnstructured](mgr.GetClient(), eventRec, status.EmitDeprecatedMetrics)

	maxConcurrentReconciles := r.cfg.GetReconcileResourceMaxConcurrency(kind)

	return ctrlrt.NewControllerManagedBy(mgr).
		For(r.rmf.ResourceDescriptor().EmptyRuntimeObject()).
		WithOptions(controller.Options{
			MaxConcurrentReconciles: maxConcurrentReconciles,
		}).
		Named(fmt.Sprintf("operatorpkg.%s.status", strings.ToLower(kind))).
		Complete(sc)
}
