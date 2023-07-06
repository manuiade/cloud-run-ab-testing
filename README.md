# A/B Testing with Cloud Run and Cloud Build

Cloudbuild based solution for performing A/B testing on Cloud Run using traffic splitting (50/50 between last and last but one)

Cloud Build revisions are updated with this procedure:

- Deploy revision 1: 100% traffic on revision 1

- Deploy revision 2: 50% traffic on revision 2 and 50% traffic on revision 1

- Deploy revision N: 50% traffic on revision N and 50% traffic on revision N - 1


## Env variables

```
PROJECT_ID=<PROJECT_ID>
REGION=europe-west1
REPO=hello-go
gcloud config set project $PROJECT_ID
```

## Create AR Repo
```
gcloud artifacts repositories create $REPO \
	--repository-format=docker \
	--location=$REGION
```

## Call Build for v1

```
gcloud builds submit ./src \
    --region $REGION \
    --config ./cloudbuild.yaml \
    --substitutions _RUN_SERVICE=hello-go,_GCP_REGION=$REGION,_PORT=8080,_ENV_VARS=FOOBAR="hello-go",_CONCURRENCY=80,_CPU=2,_MEMORY_GB=4Gi,_TIMEOUT=1000,_MIN_INSTANCES=0,_MAX_INSTANCES=1,_INGRESS_SETTINGS=all,_TAG=version1,_REPO=$REPO
```

## Modify code

## Call Build for v2

```
gcloud builds submit ./src \
    --region $REGION \
    --config ./cloudbuild.yaml \
    --substitutions _RUN_SERVICE=hello-go,_GCP_REGION=$REGION,_PORT=8080,_ENV_VARS=FOOBAR="hello-go",_CONCURRENCY=80,_CPU=2,_MEMORY_GB=4Gi,_TIMEOUT=1000,_MIN_INSTANCES=0,_MAX_INSTANCES=1,_INGRESS_SETTINGS=all,_TAG=version2,_REPO=$REPO
```

## Notice split traffic from Cloud Build between revision 1 and revision 2

## Modify code

## Call Build for v3

```
gcloud builds submit ./src \
    --region $REGION \
    --config ./cloudbuild.yaml \
    --substitutions _RUN_SERVICE=hello-go,_GCP_REGION=$REGION,_PORT=8080,_ENV_VARS=FOOBAR="hello-go",_CONCURRENCY=80,_CPU=2,_MEMORY_GB=4Gi,_TIMEOUT=1000,_MIN_INSTANCES=0,_MAX_INSTANCES=1,_INGRESS_SETTINGS=all,_TAG=version3,_REPO=$REPO
```

## Notice split traffic from Cloud Build between revision 2 and revision 3

## Split traffic between revisions using tags

```
gcloud run services update-traffic hello-go \
    --region $REGION \
    --to-tags=version1=33,version2=33,version3=34
```

## Stable on version 3

```
gcloud run services update-traffic hello-go \
    --region $REGION \
    --to-tags=version3=100
```