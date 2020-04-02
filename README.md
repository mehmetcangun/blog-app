# How to compile and run
- install docker
- open terminal
- git clone https://github.com/mehmetcangun/blog-app.git
- cd blog-app
- docker compose up
- wait until the echo server loaded.

# Issues
- Front-end
  - login state not working after refresh the page.
- Back-end
  - JWT Auth is not implemented.
  - Due to auth, posts are open for all who access the api.
