run-docker: 
			docker build --tag forum .
			docker run -p 8080:8080 -it forum