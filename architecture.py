# diagram.py
from diagrams import Cluster, Diagram
from diagrams.k8s.network import Ingress, Service
from diagrams.k8s.infra import Node
from diagrams.k8s.compute import Pod
from diagrams.k8s.storage import PersistentVolume, PersistentVolumeClaim, Volume
from diagrams.onprem.database import MySQL

with Diagram("Simple Cluster of Kubernetes", show = False):

    with Cluster("Kubernetes Cluster"):

        ing = Ingress("Ingress")
        backend_svc = Service("Backend Service")
        database_svc = Service("Database Service")
        pv = PersistentVolume("Persistent Volume")
        pvc = PersistentVolumeClaim("Persistent Volume Claim")

        ing >> backend_svc

        backend_1 = Pod("Backend-cbb496c3")
        backend_2 = Pod("Backend-ff2b4ebf")
        backend_3 = Pod("Backend-f13552aa")

        backend_svc >> [backend_1, backend_2, backend_3]
        backend_1 >> database_svc
        backend_2 >> database_svc
        backend_3 >> database_svc

        with Cluster("Pod Database"):

            volume = Volume("Volume")
            database = MySQL("Database MySQL")

            database_svc >> database
            pv >> pvc << volume
            database >> volume