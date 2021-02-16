pipeline {
  agent {
    kubernetes {
      defaultContainer 'builder'
      yamlFile 'JenkinsPod.yaml'
    }
  }
  stages {
    stage('Build docker') {
      when { tag "*" }
      steps {
        withCredentials([file(credentialsId: 'docker-hub', variable: 'DOCKER_HUB')]) {

          sh "echo \"Found TAG: ${TAG_NAME}\""

          sh "cp $DOCKER_HUB /kaniko/.docker/config.json"
          sh "[ -z \"${TAG_NAME}\" ] || /kaniko/executor --dockerfile `pwd`/Dockerfile.collider --context `pwd` --destination=nielsavonds/collider:${TAG_NAME}"
        }
      }
    }
  }
}
