package adapters

import (
	"context"

	corev1 "k8s.io/api/core/v1"
)

type PodHandler interface {
	Create(ctx context.Context, pod *corev1.Pod) error

	Update(ctx context.Context, pod *corev1.Pod) error

	Delete(ctx context.Context, pod *corev1.Pod) error

	Get(ctx context.Context, namespace, name string) (*corev1.Pod, error)

	GetStatus(ctx context.Context, namespace, name string) (*corev1.PodStatus, error)

	List(context.Context) ([]*corev1.Pod, error)

	Notify(context.Context, func(*corev1.Pod))
}

type NodeHandler interface {
	Probe(context.Context) error

	NotifyStatus(ctx context.Context, cb func(*corev1.Node))

	Configure(context.Context, *corev1.Node)
}

type PVCHandler interface {
	Get(ctx context.Context, namespace, name string) (*corev1.PersistentVolumeClaim, error)

	Delete(ctx context.Context, pvc *corev1.PersistentVolumeClaim) error

	Notify(context.Context, func(*corev1.PersistentVolumeClaim))
}

type PVHandler interface {
	Get(ctx context.Context, name string) (*corev1.PersistentVolume, error)

	Patch(ctx context.Context, old, new *corev1.PersistentVolume) (*corev1.PersistentVolume, error)

	Delete(ctx context.Context, pv *corev1.PersistentVolume) error

	Notify(context.Context, func(*corev1.PersistentVolume))
}
