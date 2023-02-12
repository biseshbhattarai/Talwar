To install and run backend : 

Install Go 

Go to https://go.dev/doc/install and click download and manage path . 

After installing go, 
cd /talwar/server 

Install these tools : 
go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest 
go install -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest 
go install -v github.com/projectdiscovery/naabu/v2/cmd/naabu@latest 

git clone https://github.com/trufflesecurity/trufflehog.git
cd trufflehog; go install

Add to /usr/local/bin from ~/go/bin 

Run the server 
go run server.go 


To install and run frontend : 
cd client/ 
npm i 
To run and start server 
npm start 