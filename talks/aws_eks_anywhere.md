## Some facts about AWS EKS Anywhere
1. No requirement of an AWS account.
2. Project is fully open sourced at: https://github.com/aws/eks-anywhere
3. AWS distro: https://github.com/aws/eks-distro
4. Bottle rocket AWS os: https://github.com/bottlerocket-os

## Demo
1. eksctl - official amazon cli for Amazon EKS.
2. eksctl anywhere - helps to bootstrap resources specific to EKS anywhere.
3. eksctl anywhere
    create - Create resources
    delete - Delete resources
    generate - Generate resources
    help - Help about any command
    upgrade - Upgrade resources
    version - Get the eksctl anywhere version
4. eksctl anywhere generate clusterconfig test -p vsphere - 'vsphere' is probably the cluster name
    gives out a bunch of configuration files (very much similar to kuberenetes config files)
    - control plane configurations
    - etcd configurations
    ability to modify things like - control plane endpoints, external etcd configuration, etc
5. similar to EKS, have declarative yaml files of configurations and use commands to create resources/objects.
6. difference is that EKS creates resources in cloud using cloudformation like tools but EKS Anywhere is specific to onprem
7. started support for only 'vcenter' (2021), plan to support arm, raspberry pi, etc in future.
8. osfamily support for bottlerocket (aws opensource api driven os), ubuntu (2021)
9. eksctl anywhere create cluster -f <filename>
    -f, --filename (filename)
        --force-cleanup (force deletion of previously created bootstrap cluster)
    -h, --help
        --skip-ip-check (skip check for whether cluster control plane ip is in use)
10. EKS Anywhere is like a wrapper around 'cluster-api'. it takes core components of cluster api for specs and crds and it takes provisioner and ties things together.
11. EKS Anywhere tries to productize open source bits like K8s cluster, etc and provide an interface to configure various configurations.
12. EKS Anywhere does not support installation and control on pre-provisioned nodes.
13. EKS Anywhere - creates new VMs, bootstraps them, etc.
14. Bare metal supports.
15. EKS Anywhere CNI - celium.
16. Git repos can be used to synchronize changes with help of Gitops. - Gitops inside the cluster recongnizes the changes by a pull based method instead of something like Jenkins/ArgoCD doing push based model.
17. eksctl upgrade use case
    change the configurations in the file.
    eksctl anywhere upgrade -f <filepath>
18. https://github.com/aws/eks-anywhere/tree/main/docs
19. EKS anywhere + Tinkerbell - https://github.com/aws/eks-anywhere/tree/main/examples/tinkerbell
20. 
    
    