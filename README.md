# GCP Light Viewer

# About

GCP Light Viewer show some informations about your projects and can help you to see information on Clusters.
Now you can check on web, and don't need applications just like Lens, or browse on slow gcp web console to get some information.


### Overview
You can list project and Clusters with this application

    ➜  gcp-golang git:(main) go run main.go
    statusCode: 200
    Project[0]: lightbank-333002
    statusCode: 200
    Cluster[0]: cluster-lightbank
        Network[0]: default
        NetworkConfig[0]: projects/lightbank-333002/global/networks/default
        NetworkConfig[0]: projects/lightbank-333002/regions/us-central1/subnetworks/default 