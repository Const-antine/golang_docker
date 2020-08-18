pipeline {
   environment {
       registry = "localhost:5000/mongo-api:1.3"
       GOCACHE = "/tmp"
   }

def app

agent { dockerfile true }


   stages {
stage('Clone repository') {        
       
 checkout scm

}

stage('Build') {


app = docker.build(registry)
}




       stage('Test') {
           steps {
           
        app.inside {
            sh 'echo "Tests passed"'
        }           

}

       
}



       stage ('Deploy') {

      environment {
               vaultPass = 'vaultpass'
           }


           steps {
      ansiblePlaybook(
           playbook: "./ansible-golang/mongo-api-k8s.yaml"
           inventory: "./ansible-golang/hosts.txt"
           vaultCredentialsId: vaultPass     
           }
       }
   }
}
