#!/bin/bash

echo "Creando namespace..."
kubectl apply -f namespace.yaml

echo "Creando Persistent Volume y Claim..."
kubectl apply -f pv.yaml
kubectl apply -f pvc.yaml

echo "Desplegando MongoDB..."
kubectl apply -f mongo-deployment.yaml
kubectl apply -f mongo-service.yaml

echo "Desplegando microservicio p-go-delete..."
kubectl apply -f delete-deployment.yaml
kubectl apply -f delete-service.yaml

echo "Configurando Ingress..."
kubectl apply -f ingress.yaml
