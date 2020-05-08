docker build --no-cache -t emailtovamos/readinessprobe:v1 .

docker push emailtovamos/readinessprobe:v1

kubectl delete deployment readinessprobe

kubectl apply -f devops/deployment.yaml