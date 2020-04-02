# How to compile and run
- install docker compose
- open terminal
- git clone https://github.com/mehmetcangun/blog-app.git
- cd blog-app
- docker-compose up
- wait until the echo server loaded.
- click on the 2nd url

![Simple simulation](https://imgur.com/R7w9dvl)

# Issues
- Front-end
  - login state not working after refresh the page.
- Back-end
  - JWT Auth is not implemented.
  - Due to auth, posts are open for all who access the api.
