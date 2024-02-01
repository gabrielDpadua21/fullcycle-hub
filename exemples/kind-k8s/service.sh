#!/bin/bash

## create cluster
kind create cluster --config=kind.yaml --name=fullcycle

## add cluster in context
kubectl cluster-info --context kind-fullcycle

## up pods
kubectl apply -f deployment.yaml

## up services
kubectl apply -f service-loadbalancer.yaml

