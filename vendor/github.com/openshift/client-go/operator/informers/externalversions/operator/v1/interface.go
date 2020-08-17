// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/openshift/client-go/operator/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Authentications returns a AuthenticationInformer.
	Authentications() AuthenticationInformer
	// CSISnapshotControllers returns a CSISnapshotControllerInformer.
	CSISnapshotControllers() CSISnapshotControllerInformer
	// Configs returns a ConfigInformer.
	Configs() ConfigInformer
	// Consoles returns a ConsoleInformer.
	Consoles() ConsoleInformer
	// DNSes returns a DNSInformer.
	DNSes() DNSInformer
	// Etcds returns a EtcdInformer.
	Etcds() EtcdInformer
	// IngressControllers returns a IngressControllerInformer.
	IngressControllers() IngressControllerInformer
	// KubeAPIServers returns a KubeAPIServerInformer.
	KubeAPIServers() KubeAPIServerInformer
	// KubeControllerManagers returns a KubeControllerManagerInformer.
	KubeControllerManagers() KubeControllerManagerInformer
	// KubeSchedulers returns a KubeSchedulerInformer.
	KubeSchedulers() KubeSchedulerInformer
	// KubeStorageVersionMigrators returns a KubeStorageVersionMigratorInformer.
	KubeStorageVersionMigrators() KubeStorageVersionMigratorInformer
	// Networks returns a NetworkInformer.
	Networks() NetworkInformer
	// OpenShiftAPIServers returns a OpenShiftAPIServerInformer.
	OpenShiftAPIServers() OpenShiftAPIServerInformer
	// OpenShiftControllerManagers returns a OpenShiftControllerManagerInformer.
	OpenShiftControllerManagers() OpenShiftControllerManagerInformer
	// ServiceCAs returns a ServiceCAInformer.
	ServiceCAs() ServiceCAInformer
	// ServiceCatalogAPIServers returns a ServiceCatalogAPIServerInformer.
	ServiceCatalogAPIServers() ServiceCatalogAPIServerInformer
	// ServiceCatalogControllerManagers returns a ServiceCatalogControllerManagerInformer.
	ServiceCatalogControllerManagers() ServiceCatalogControllerManagerInformer
	// Storages returns a StorageInformer.
	Storages() StorageInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Authentications returns a AuthenticationInformer.
func (v *version) Authentications() AuthenticationInformer {
	return &authenticationInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CSISnapshotControllers returns a CSISnapshotControllerInformer.
func (v *version) CSISnapshotControllers() CSISnapshotControllerInformer {
	return &cSISnapshotControllerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Configs returns a ConfigInformer.
func (v *version) Configs() ConfigInformer {
	return &configInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Consoles returns a ConsoleInformer.
func (v *version) Consoles() ConsoleInformer {
	return &consoleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// DNSes returns a DNSInformer.
func (v *version) DNSes() DNSInformer {
	return &dNSInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Etcds returns a EtcdInformer.
func (v *version) Etcds() EtcdInformer {
	return &etcdInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// IngressControllers returns a IngressControllerInformer.
func (v *version) IngressControllers() IngressControllerInformer {
	return &ingressControllerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// KubeAPIServers returns a KubeAPIServerInformer.
func (v *version) KubeAPIServers() KubeAPIServerInformer {
	return &kubeAPIServerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// KubeControllerManagers returns a KubeControllerManagerInformer.
func (v *version) KubeControllerManagers() KubeControllerManagerInformer {
	return &kubeControllerManagerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// KubeSchedulers returns a KubeSchedulerInformer.
func (v *version) KubeSchedulers() KubeSchedulerInformer {
	return &kubeSchedulerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// KubeStorageVersionMigrators returns a KubeStorageVersionMigratorInformer.
func (v *version) KubeStorageVersionMigrators() KubeStorageVersionMigratorInformer {
	return &kubeStorageVersionMigratorInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Networks returns a NetworkInformer.
func (v *version) Networks() NetworkInformer {
	return &networkInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// OpenShiftAPIServers returns a OpenShiftAPIServerInformer.
func (v *version) OpenShiftAPIServers() OpenShiftAPIServerInformer {
	return &openShiftAPIServerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// OpenShiftControllerManagers returns a OpenShiftControllerManagerInformer.
func (v *version) OpenShiftControllerManagers() OpenShiftControllerManagerInformer {
	return &openShiftControllerManagerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ServiceCAs returns a ServiceCAInformer.
func (v *version) ServiceCAs() ServiceCAInformer {
	return &serviceCAInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ServiceCatalogAPIServers returns a ServiceCatalogAPIServerInformer.
func (v *version) ServiceCatalogAPIServers() ServiceCatalogAPIServerInformer {
	return &serviceCatalogAPIServerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ServiceCatalogControllerManagers returns a ServiceCatalogControllerManagerInformer.
func (v *version) ServiceCatalogControllerManagers() ServiceCatalogControllerManagerInformer {
	return &serviceCatalogControllerManagerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Storages returns a StorageInformer.
func (v *version) Storages() StorageInformer {
	return &storageInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
