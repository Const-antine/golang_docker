pipeline {
   agent any
   environment {
       registry = "localhost:5000/mongo-api:1.3"
       GOCACHE = "/tmp"
      
   }


   stages {
       stage('Clone repository') {

 steps {
     script {
     checkout scm
     }
}
}
      
      
      stage ('Test') {
         steps {
            sh "ls -la"
            sh "pwd"
            sh "ls ./ansible-golang/mongo-api-k8s.yaml"
         }
      }

       stage ('Deploy') {

      environment {
               vaultPass = 'vaultpass'
           }


           steps {
      ansiblePlaybook(
           playbook: "./ansible-golang/mongo-api-k8s.yaml",
           inventory: "./ansible-golang/hosts.txt",
           vaultCredentialsId: vaultPass  
           )
           }
       }
   }
}

