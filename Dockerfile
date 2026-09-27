FROM golang:1.27-bookworm

# Add 'git' to the installation list
RUN apt-get update && apt-get install -y python3 python3-pip git

WORKDIR /app
COPY . .

# Clone your existing GitHub repository into the container
# Replace the URL with your actual repository link
RUN git clone https://github.com/JacksonS25/Song_Set_Maker.git ./projects/Song_Set_Maker

# IF your Python project has a requirements.txt, uncomment the line below to install dependencies. 
# (Debian 12 requires the --break-system-packages flag for global Docker pip installs)
# RUN pip3 install -r ./projects/Song_Set_Maker/requirements.txt --break-system-packages

RUN go build -o server .

EXPOSE 3000
CMD ["./server"]