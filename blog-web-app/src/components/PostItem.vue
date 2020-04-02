<template>
    <div class="section">
        <div class="notification has-background-white-bis">
            <div class="columns has-text-centered is-vcentered" style="margin:0;padding:0">
                <span class="is-pulled-right column is-2 is-family-monospace is-italic is-lowercase">{{ post.estimatedTime }} minutes to read.</span>
                <hr width="100%">
                <span class="is-pulled-right column is-2 is-family-monospace is-italic is-uppercase">{{ createdDate }}</span>
            </div>
            <button 
                class="button modal-button is-dark is-outlined is-medium is-fullwidth" 
                @click="isComponentModalActive = true"
            >
                <strong>{{post.title}}</strong>
            </button>
            <b-modal 
                :active.sync="isComponentModalActive"
                scroll="keep"
                >
                <div class="card">
                    <div class="card-image">
                        <figure class="image is-3by1">
                            <img :src=post.imgSrc />
                        </figure>
                    </div>
                    <div class="card-content">
                        <button class="modal-close is-large" @click="$parent.close()"></button>
                        <h2 class="title is-2">{{ post.title }}</h2>
                        <hr>
                        <div class="section">
                            <p class="content">{{ post.content }}</p>
                        </div>
                        <div class="columns is-vcentered has-background-white-bis" id="profile">
                            <div class="column is-two-thirds">
                                <h2 class="sub-title is-2 is-italic">Author Details</h2>
                            </div>
                            <div class="column has-text-right">
                                {{ post.userDetails.name }} <span class="is-uppercase has-text-weight-bold"> {{ post.userDetails.surname }}</span>
                            </div>
                            <div class="column">
                                <figure class="image is-128x128">
                                    <img class="is-rounded" :src=post.userDetails.imgSrc>
                                </figure>
                            </div>
                        </div>
                    </div>
                </div>
            </b-modal> 
        </div>
    </div>
</template>


<script>
    export default {
        name: 'PostItem',
        props: {
            post: Object
        },
        data(){
            return {
                isComponentModalActive: false
            }
        },
        computed: {
            createdDate: function(){
                const date = new Date(this.post.createdDate)
                
                return date.toLocaleDateString(undefined, {
                    month: 'long',
                    year: 'numeric',
                    day: 'numeric'
                })
            }
        }
    }
</script>