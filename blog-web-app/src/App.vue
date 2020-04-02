<template>
  <div id="app" class="section">
    <div class="container">
      <div>
          <div v-if="isLoggedIn">
              <div class="columns is-vcentered" id="profile">
                  <div class="column is-two-thirds">
                      <h2 class="title is-2">The Posts</h2>
                  </div>
                  
                  <div class="column has-text-right">
                      {{User.name}} <span class="is-uppercase has-text-weight-bold">{{User.surname}}</span>
                      <br>
                      <b-button 
                          class="is-text is-italic"
                          @click="logOut"
                      >Logout</b-button> 
                  </div>
                  <div class="column">
                      <figure class="image is-128x128">
                          <img class="is-rounded" :src=User.imgSrc>
                      </figure>
                  </div>
              </div>
              <hr>
              <b-button 
                  class="button is-primary is-medium is-fullwidth"
                  @click="isModalActivePost = true"
              >New Post!</b-button>
              <b-modal
                  :active.sync="isModalActivePost"
                  has-modal-card
                  trap-focus
                  aria-role="dialog"
                  aria-modal
              >
                  <div class="modal-card">
                      <header class="modal-card-head">
                          <h2 class="is-title is-2">New Post</h2>
                      </header>
                      <section class="modal-card-body">
                          <form action="">
                              <b-field label="Title">
                                  <b-input v-model="title" @input="emptyness" required></b-input>
                              </b-field>
                              <b-field label="Content">
                                  <b-input v-model="content" @input="emptyness" type="textarea" required></b-input>
                              </b-field>
                              <b-field label="Image source">
                                  <b-input v-model="imgSrc" @input="emptyness" type="url" required></b-input>
                              </b-field>
                              <b-field label="Read time">
                                  <b-numberinput v-model="estimatedTime" @input="emptyness" min="0" required></b-numberinput>
                              </b-field>
                              <b-button :disabled="empty"
                                  class="is-primary is-fullwidth"
                                  @click="newPost">New Post!</b-button>
                          </form>
                      </section>
                  </div>
              </b-modal>
          </div>
          <div v-else class="has-text-centered">
              <h2 class="title is-2"> Home Page </h2>
              <button 
                  class="button is-primary is-medium is-fullwidth" 
                  @click="isModalActive = true"
              >Login!</button>
              <b-modal 
                  :active.sync="isModalActive"
                  scroll="keep"
                  has-modal-card
                  trap-focus
                  aria-role="dialog"
                  aria-modal
                  >
                  <div class="modal-card">
                      <header class="modal-card-head">
                          <h1 class="modal-card-title">Login</h1>
                          <button class="modal-close is-large" @click="$parent.close()"></button>
                      </header>
                      <section class="modal-card-body">
                          <form>
                              <div class="block">
                                  <b-field label="Username" native-value="on-border">
                                      <b-input v-model="username" required></b-input>
                                  </b-field>
                              </div>
                              <b-field label="Password">
                                  <b-input v-model="password" type="password" required password-reveal></b-input>
                              </b-field>
                              <b-button class="is-dark" @click="checkLogin">Login!</b-button>
                          </form>
                      </section>
                  </div>
              </b-modal>
          </div>
      </div>
      <PostItem 
        v-for="post in posts"
        v-bind:post="post"
        v-bind:key="post.id"
      ></PostItem>
    </div>
  </div>
</template>

<script>
import PostItem from './components/PostItem.vue'

export default {
  name: 'App',
  components: {
    PostItem
  },
  mounted() {
    this.getPosts()
  },
  data() {
    return {
      posts: [],
      isModalActive: false,
      isModalActivePost: false,
      username: "root",
      password: "toor",
      User: null,
      title: "",
      content: "",
      imgSrc: "",
      estimatedTime: 0,
      empty: true,
      isLoggedIn: false
    }
  },
  methods: {
    isSuccess(msg){
        this.$buefy.notification.open({
            duration: 1000,
            message: msg,
            position: 'is-bottom-right',
            type: 'is-success',
            hasIcon: true
        })
    },
    isWrong(msg){
        this.$buefy.notification.open({
            duration: 1000,
            message: msg,
            position: 'is-bottom-right',
            type: 'is-danger',
            hasIcon: true
        })
    },
    async getPosts() {
      try {
        const response = await this.$axios.get("posts")
        this.posts = response.data.posts
        this.isSuccess("Posts are completed.")
      } catch {
        this.isWrong("There is an issue!")
      }
    },
    async checkLogin() {
        try {
            await this.$axios.post('login', {
                username: this.username,
                password: this.password
            })
            .then(response => {

                this.User = response.data
                this.isLoggedIn = true
                this.isSuccess("Login succeed")
                this.isModalActive = false
            })    
        } catch {
            this.User = null
            this.isLoggedIn = false
            this.isWrong("Username or Password wrong!")
        }
    },
    logOut(){
        this.User = null
        this.isLoggedIn = false
    },
    async newPost(){
        try {
            await this.$axios.post('posts', {
                "title": this.title,
                "content": this.content,
                "imgSrc": this.imgSrc,
                "estimatedTime": this.estimatedTime,
                "userDetails": this.User
            })
            .then(response => {
                this.posts.push({
                  "id": response.affected,
                  "title": this.title,
                  "content": this.content,
                  "imgSrc": this.imgSrc,
                  "createdDate": Date.now(),
                  "estimatedTime": this.estimatedTime,
                  "userDetails": this.User
                })
                this.posts.reverse()
                this.isSuccess("New post is created.")
                this.isModalActivePost = false
            })
        } catch {
            this.isWrong("There is an issue!")
        }
    }, 
    emptyness() {
        if( this.title === "" || this.content === "" || this.imgSrc === "" || this.estimatedTime === 0)
            this.empty = true
        else
            this.empty = false
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
.section {
    padding: 1.5rem;
}

hr {
    margin: 1rem 0
}
</style>
