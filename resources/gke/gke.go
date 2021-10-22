package gke

import (
	"encoding/json"
	"fmt"
	"gcp-golang/pkg/web"
	"log"
	"time"
)

type AllClusters struct {
	Clusters []struct {
		Name       string `json:"name"`
		NodeConfig struct {
			MachineType string   `json:"machineType"`
			DiskSizeGb  int      `json:"diskSizeGb"`
			OauthScopes []string `json:"oauthScopes"`
			Metadata    struct {
				DisableLegacyEndpoints string `json:"disable-legacy-endpoints"`
			} `json:"metadata"`
			ImageType              string `json:"imageType"`
			ServiceAccount         string `json:"serviceAccount"`
			DiskType               string `json:"diskType"`
			WorkloadMetadataConfig struct {
				NodeMetadata string `json:"nodeMetadata"`
				Mode         string `json:"mode"`
			} `json:"workloadMetadataConfig"`
			ShieldedInstanceConfig struct {
				EnableSecureBoot          bool `json:"enableSecureBoot"`
				EnableIntegrityMonitoring bool `json:"enableIntegrityMonitoring"`
			} `json:"shieldedInstanceConfig"`
		} `json:"nodeConfig"`
		MasterAuth struct {
			ClusterCaCertificate string `json:"clusterCaCertificate"`
		} `json:"masterAuth"`
		LoggingService    string `json:"loggingService"`
		MonitoringService string `json:"monitoringService"`
		Network           string `json:"network"`
		ClusterIpv4Cidr   string `json:"clusterIpv4Cidr"`
		AddonsConfig      struct {
			HTTPLoadBalancing struct {
			} `json:"httpLoadBalancing"`
			HorizontalPodAutoscaling struct {
			} `json:"horizontalPodAutoscaling"`
			KubernetesDashboard struct {
				Disabled bool `json:"disabled"`
			} `json:"kubernetesDashboard"`
			NetworkPolicyConfig struct {
				Disabled bool `json:"disabled"`
			} `json:"networkPolicyConfig"`
			DNSCacheConfig struct {
				Enabled bool `json:"enabled"`
			} `json:"dnsCacheConfig"`
			GcePersistentDiskCsiDriverConfig struct {
				Enabled bool `json:"enabled"`
			} `json:"gcePersistentDiskCsiDriverConfig"`
		} `json:"addonsConfig"`
		Subnetwork string `json:"subnetwork"`
		NodePools  []struct {
			Name   string `json:"name"`
			Config struct {
				MachineType string   `json:"machineType"`
				DiskSizeGb  int      `json:"diskSizeGb"`
				OauthScopes []string `json:"oauthScopes"`
				Metadata    struct {
					DisableLegacyEndpoints string `json:"disable-legacy-endpoints"`
				} `json:"metadata"`
				ImageType              string `json:"imageType"`
				ServiceAccount         string `json:"serviceAccount"`
				DiskType               string `json:"diskType"`
				WorkloadMetadataConfig struct {
					NodeMetadata string `json:"nodeMetadata"`
					Mode         string `json:"mode"`
				} `json:"workloadMetadataConfig"`
				ShieldedInstanceConfig struct {
					EnableSecureBoot          bool `json:"enableSecureBoot"`
					EnableIntegrityMonitoring bool `json:"enableIntegrityMonitoring"`
				} `json:"shieldedInstanceConfig"`
			} `json:"config"`
			InitialNodeCount int `json:"initialNodeCount"`
			Autoscaling      struct {
				Enabled         bool `json:"enabled"`
				MaxNodeCount    int  `json:"maxNodeCount"`
				Autoprovisioned bool `json:"autoprovisioned"`
			} `json:"autoscaling"`
			Management struct {
				AutoUpgrade bool `json:"autoUpgrade"`
				AutoRepair  bool `json:"autoRepair"`
			} `json:"management"`
			MaxPodsConstraint struct {
				MaxPodsPerNode string `json:"maxPodsPerNode"`
			} `json:"maxPodsConstraint"`
			PodIpv4CidrSize int      `json:"podIpv4CidrSize"`
			Locations       []string `json:"locations"`
			NetworkConfig   struct {
				PodRange         string `json:"podRange"`
				PodIpv4CidrBlock string `json:"podIpv4CidrBlock"`
			} `json:"networkConfig"`
			SelfLink          string   `json:"selfLink"`
			Version           string   `json:"version"`
			InstanceGroupUrls []string `json:"instanceGroupUrls"`
			Status            string   `json:"status"`
			UpgradeSettings   struct {
				MaxSurge int `json:"maxSurge"`
			} `json:"upgradeSettings"`
		} `json:"nodePools"`
		Locations        []string `json:"locations"`
		LabelFingerprint string   `json:"labelFingerprint"`
		LegacyAbac       struct {
		} `json:"legacyAbac"`
		IPAllocationPolicy struct {
			UseIPAliases               bool   `json:"useIpAliases"`
			ClusterIpv4Cidr            string `json:"clusterIpv4Cidr"`
			ServicesIpv4Cidr           string `json:"servicesIpv4Cidr"`
			ClusterSecondaryRangeName  string `json:"clusterSecondaryRangeName"`
			ServicesSecondaryRangeName string `json:"servicesSecondaryRangeName"`
			ClusterIpv4CidrBlock       string `json:"clusterIpv4CidrBlock"`
			ServicesIpv4CidrBlock      string `json:"servicesIpv4CidrBlock"`
		} `json:"ipAllocationPolicy"`
		MasterAuthorizedNetworksConfig struct {
		} `json:"masterAuthorizedNetworksConfig"`
		MaintenancePolicy struct {
			ResourceVersion string `json:"resourceVersion"`
		} `json:"maintenancePolicy"`
		Autoscaling struct {
			EnableNodeAutoprovisioning bool `json:"enableNodeAutoprovisioning"`
			ResourceLimits             []struct {
				ResourceType string `json:"resourceType"`
				Maximum      string `json:"maximum"`
			} `json:"resourceLimits"`
			AutoscalingProfile               string `json:"autoscalingProfile"`
			AutoprovisioningNodePoolDefaults struct {
				OauthScopes     []string `json:"oauthScopes"`
				ServiceAccount  string   `json:"serviceAccount"`
				UpgradeSettings struct {
					MaxSurge int `json:"maxSurge"`
				} `json:"upgradeSettings"`
				Management struct {
					AutoUpgrade bool `json:"autoUpgrade"`
					AutoRepair  bool `json:"autoRepair"`
				} `json:"management"`
				ImageType string `json:"imageType"`
			} `json:"autoprovisioningNodePoolDefaults"`
		} `json:"autoscaling"`
		NetworkConfig struct {
			Network                   string `json:"network"`
			Subnetwork                string `json:"subnetwork"`
			EnableIntraNodeVisibility bool   `json:"enableIntraNodeVisibility"`
			DefaultSnatStatus         struct {
			} `json:"defaultSnatStatus"`
			ServiceExternalIpsConfig struct {
				Enabled bool `json:"enabled"`
			} `json:"serviceExternalIpsConfig"`
		} `json:"networkConfig"`
		DefaultMaxPodsConstraint struct {
			MaxPodsPerNode string `json:"maxPodsPerNode"`
		} `json:"defaultMaxPodsConstraint"`
		AuthenticatorGroupsConfig struct {
		} `json:"authenticatorGroupsConfig"`
		DatabaseEncryption struct {
			State string `json:"state"`
		} `json:"databaseEncryption"`
		VerticalPodAutoscaling struct {
			Enabled bool `json:"enabled"`
		} `json:"verticalPodAutoscaling"`
		ShieldedNodes struct {
			Enabled bool `json:"enabled"`
		} `json:"shieldedNodes"`
		ReleaseChannel struct {
			Channel string `json:"channel"`
		} `json:"releaseChannel"`
		WorkloadIdentityConfig struct {
			IdentityNamespace string `json:"identityNamespace"`
			WorkloadPool      string `json:"workloadPool"`
			IdentityProvider  string `json:"identityProvider"`
		} `json:"workloadIdentityConfig"`
		ClusterTelemetry struct {
			Type string `json:"type"`
		} `json:"clusterTelemetry"`
		NotificationConfig struct {
			Pubsub struct {
			} `json:"pubsub"`
		} `json:"notificationConfig"`
		SelfLink              string    `json:"selfLink"`
		Zone                  string    `json:"zone"`
		Endpoint              string    `json:"endpoint"`
		InitialClusterVersion string    `json:"initialClusterVersion"`
		CurrentMasterVersion  string    `json:"currentMasterVersion"`
		CurrentNodeVersion    string    `json:"currentNodeVersion"`
		CreateTime            time.Time `json:"createTime"`
		Status                string    `json:"status"`
		ServicesIpv4Cidr      string    `json:"servicesIpv4Cidr"`
		InstanceGroupUrls     []string  `json:"instanceGroupUrls"`
		Location              string    `json:"location"`
		Master                struct {
		} `json:"master"`
		Autopilot struct {
			Enabled bool `json:"enabled"`
		} `json:"autopilot"`
		ID               string `json:"id"`
		NodePoolDefaults struct {
			NodeConfigDefaults struct {
			} `json:"nodeConfigDefaults"`
		} `json:"nodePoolDefaults"`
		LoggingConfig struct {
			ComponentConfig struct {
				EnableComponents []string `json:"enableComponents"`
			} `json:"componentConfig"`
		} `json:"loggingConfig"`
		MonitoringConfig struct {
			ComponentConfig struct {
				EnableComponents []string `json:"enableComponents"`
			} `json:"componentConfig"`
		} `json:"monitoringConfig"`
	} `json:"clusters"`
}

func Get() {
	//url := fmt.Sprintf("projects/%s/locations/-", "ultra-sound-324019")
	url := "https://container.googleapis.com/v1beta1/projects/ultra-sound-324019/locations/us-central1/clusters"

	data, err := web.Get(url)
	if err != nil {
		log.Println("Clusters:", err)
	}

	var clusters AllClusters
	json.Unmarshal(data, &clusters)
	for i, v := range clusters.Clusters {
		fmt.Printf("Cluster[%d]: %s \n", i, v.Name)
		fmt.Printf("\tNetwork[%d]: %s \n", i, v.Network)
		fmt.Printf("\tNetworkConfig[%d]: %s \n", i, v.NetworkConfig.Network)
		fmt.Printf("\tNetworkConfig[%d]: %s \n", i, v.NetworkConfig.Subnetwork)
	}
}
