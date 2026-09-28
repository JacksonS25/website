FROM golang:1.27-bookworm

# Add 'git' to the installation list
RUN apt-get update && apt-get install -y python3 python3-pip git

WORKDIR /app
COPY . .

# Clone your Song_Set_Maker GitHub repository into the container
RUN git clone https://github.com/JacksonS25/Song_Set_Maker.git ./projects/Song_Set_Maker

# Install all dependencies from the requirements file
RUN pip3 install -r ./projects/Song_Set_Maker/requirements.txt --break-system-packages

RUN go build -o server .

EXPOSE 3000
CMD ["./server"]