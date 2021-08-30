Name: Matheus Arruda Ferreira de Santana.

### Prerequisites (Install the latest git package, docker, packer and Jenkins)

#### Install Latest Git Package

- Check if you have git on your machine. To check, first type the key combination Ctrl + Alt + T and execute the command: `git --version`;
- If nothing appears after you execute the above command, install git on your machine. To do that, on the terminal, type the command `sudo apt install git-all`;
- Now, check the install of git by running: `git --version`;
- You need to clone this repository.

#### Install Docker:

- To install docker, go to the terminal and type the following commands:
- `sudo apt-get update`
- `sudo apt-get install docker-ce docker-ce-cli containerd.io`
- Now, verify if docker is correctly installed by running "Hello-world" image typing the command: sudo docker run hello-world; If the following message appears, it means that docker installation appears to be working correctly.
```
Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.
```
To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/

 ##### Post-Installations:

 **Manage Docker as a non-root user:**

- First, you need to create a docker group. Go to the terminal and run: `sudo groupadd docker`;
- Add a user to the docker group by running: `sudo usermod -aG docker $USER`;
- Now, Log out and log back your machine.
- After that, back to the terminal and run the command: `newgrp docker` ;
- Now try to run a docker command without sudo by running: `docker run hello-world`. If a informational message appears, means that it works.

#### Create a Dockerhub account:

After that, you need to create a Dockerhub account. To do that, please click [here](https://hub.docker.com).
Then, create your public repository.

#### Install Packer:

- To install packer, on the terminal enter the command: `sudo apt install packer`;
- Now, check the install of git by running: `packer --version`.

#### Install Jenkins:

- To install jenkins, on the terminal enter the following commands:

```
wget -q -O - https://pkg.jenkins.io/debian-stable/jenkins.io.key | sudo apt-key add -
sudo sh -c 'echo deb https://pkg.jenkins.io/debian-stable binary/ > \
    /etc/apt/sources.list.d/jenkins.list'
sudo apt-get update
sudo apt-get install jenkins
```

##### Start Jenkins:

- Type the command: `sudo systemctl daemon-reload`;
- After that, you can start jenkins with the command: `sudo systemctl start jenkins`;
- You can check the status of the Jenkins service using the command: `sudo systemctl status jenkins`;

##### Unlock Jenkins:

- To unlock and access jenkins, go to the browser and enter the url: `http://localhost:8080` and waits until unlock jenkins page appears;
- To put the password, on terminal type the command: `sudo cat /var/lib/jenkins/secrets/initialAdminPassword`. This command will print the password at console. Copy the password and paste on the unlock jenkins page and click on continue.
- After unlocking Jenkins, the Customize Jenkins page appears. Click on the "Install suggested plugins" option.
- When the Create First Admin User page appears, specify the details for your administrator user in the respective fields and click Save and Finish. 


Homework Jenkins
================

1. Create a 2 PIPELINEs for your GO Microservice
   1.1 First pipeline will be caled BAKE and you will need use PACKER and bake a docker image
   1.2 Second pipeline will be called LAUNCH and you will DEPLOY you microservice in DOCKER

### Pipeline:

#### Job1 - Bake:

Create a new job on Jenkins, put the name "Bake" for the job and click in the option: Pipeline.
On the tab "Advanced Project Options", go to "Definition" and select the option "Pipeline Script from SCM".
After that, on "SCM" select the option "Git". Then paste this link: https://github.com/MatheusSantana95/tema12.
On the option "Branch Specifier (blank for 'any')", type */main.
On the option "Script Path", type `bake/Jenkinsfile`
Finally, click on save.

#### Job2 - Launch:
Create a new job on Jenkins, put the name "Launch" for the job and click in the option: Pipeline.
On the tab "Advanced Project Options", go to "Definition" and select the option "Pipeline Script from SCM".
After that, on "SCM" select the option "Git". Then paste this link: https://github.com/MatheusSantana95/tema12.
On the option "Branch Specifier (blank for 'any')", type `*/main`.
On the option "Script Path", type `launch/Jenkinsfile`
Finally, click on save.

#### Building the jobs:

Go to your jenkins "Dashboard" and select the first job that you´ve created and select the option "Build now".
Wait until the proccess finish.
If you receive the message "SUCCESS", your job was build successfully.
If you received the message "FAILURE", please repeat all the proccess of this tutorial carefully.
Go to your jenkins "Dashboard" and select the second job that you´ve created and select the option "Build now".
Wait until the proccess finish.
If you receive the message "SUCCESS", your job was build successfully.
If you received the message "FAILURE", please repeat all the proccess of this tutorial carefully.

### How to use the Application:

- This application is on the port: 8083. You can do the sum, subtract, multiply and division opeations accessing the endpoint as below:

```
http://localhost:8083/calc/sum/$a/$b 
http://localhost:8083/calc/sub/$a/$b 
http://localhost:8083/calc/mul/$a/$b 
http://localhost:8083/calc/div/$a/$b
```

Note: The paremeters $a and $b is where you´ll put the desired numbers to do the desired operation.

- To access the historic, you can access the endpoint: `http://localhost:8083/calc/history`

- To stop your go microservice first type the command: `docker ps`;
- Copy the container id; 
- Then type the command: `docker stop {YOUR_CONTAINER_ID}`.