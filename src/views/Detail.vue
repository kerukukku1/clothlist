<template>
    <div>
        <div class="head-spacer"></div>
        <div class="title-content">
            Detail - {{mydetail.Title}}
        </div>
        <div :style="boximage">
            <div class="image-wrapper" @click="open=true">
                click to fullscreen
            </div>
        </div>
        <image-modal v-if="open" v-on:close="open=false">
            <img :class="{'w-image' : longW, 'h-image' : !longW}" :src="require('@/server/'+this.path)"/>
        </image-modal>
        <div class="markdown-gridbox">
            <transition name="marked" out-in>
                <textarea v-if="isEdit" class="editor" v-model="content"></textarea>
            </transition>
            <div v-html="compileMarkdown"></div>
        </div>
        <div class="edit-button" @click="actEdit">{{buttonText}}</div>
        <div class="back-button">Back Home</div>
    </div>
</template>

<script lang="ts">
import {Vue, Component, Watch} from 'vue-property-decorator'
import axios from 'axios'
import modal from '@/components/ImageModal.vue'
import { Marked } from 'marked-ts';

interface boxImage{
    width : string,
    height : string,
    minHeight : string,
    backgroundSize : string,
    backgroundPosition : string,
    "backgroundImage" : string,
    margin : string
}

@Component({
    components: {
        'image-modal' : modal
    }
})
export default class Detail extends Vue {
    id : string = ""
    open : boolean = false;
    mydetail : string = "aeiuo"
    path : string = ""
    longW : boolean = false
    content : string = ""
    isEdit : boolean = false
    buttonContext : string = ""
    boximage: boxImage = {
        width : "100vw",
        height : "auto",
        minHeight : "50vh",
        backgroundSize : "cover",
        backgroundPosition : "center",
        'backgroundImage' : '',
        margin : "0 auto"
    }

    get buttonText() : string {
        return (this.isEdit)?"Save Markdown" : "Edit Markdown"
    }

    get compileMarkdown() : string{
        return Marked.parse(this.content);
    }

    actEdit(){
        this.isEdit=!this.isEdit
        if(this.isEdit)return
        let postData = new FormData()
        postData.append('target', this.mydetail.DetailID)
        postData.append('content', this.content)
        const options = {
            headers : {
                'content-type': 'application/json',
            }
        }
        axios.post('http://localhost:5000/detail/post', postData, options)
        .then(function (res){
            console.log(res)
        })
        .catch(function (err){
            console.log(err)
        })
    }

    getDetailData() {
        axios.get('http://localhost:5000/api/detail/'+this.mydetail.DetailID, {
            headers : {
                'Content-type' : 'application-json'
            }
        }).then(function (res : any){
            this.content = res.data.Comment
            console.log(res)
        }.bind(this)).catch(function (err){
            console.log(err)
        })
    }

    created() {
        this.id = this.$route.params.id
        axios.get('http://localhost:5000/api/images/'+this.id, {
            headers : {
                'Content-type' : 'application-json'
            }
        }).then(function (res){
            console.log(res.data)
            this.mydetail = res.data
            this.path = this.mydetail.Path
            Vue.set(this.boximage, 'backgroundImage', 'url('+require('@/server/'+this.path)+')')
            let img = new Image
            img.src = require('@/server/'+this.path);
            img.onload = function() {
                if (img.naturalWidth > img.naturalHeight){
                    this.longW = true
                }else{
                    this.longW = false
                }
            }.bind(this)
            this.getDetailData()
        }.bind(this)).catch(function (err){
            console.log(err)
        })
    }
}
</script>

<style scoped>
.image-wrapper {
    z-index: 100;
    width: 100vw;
    height: 50vh;
    background: rgba(0, 0, 0, 0.3);
    color: white;
    font-weight: bold;
    font-size: 40px;
    display: flex;
    justify-content: center;
    align-items: center;
    text-decoration-line: underline;
    cursor: pointer;
}

.markdown-gridbox{
    display: grid;
    grid-gap: 10px;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
}

.w-image {
    width: 80vw;
    height: auto;
    max-height: 90vh;
}
.h-image {
    width: auto;
    height: 100vh;
}

.edit-button{
    background-color:rgb(69, 196, 69);
    color: white;
    display: inline-block;
    font-size: 16px;
    padding: 10px 30px 10px 30px;
    border: none;
    text-align: center;
    font-weight: bold;
    margin: 10px;
    cursor: pointer;
}

.edit-button:hover{
    background-color:rgb(65, 209, 65);
    transition: 300ms;
}

.back-button{
    background-color:rgb(142, 142, 142);
    color: white;
    display: inline-block;
    font-size: 16px;
    padding: 10px 30px 10px 30px;
    border: none;
    text-align: center;
    font-weight: bold;
    margin: 10px;
    cursor: pointer;
}

.back-button:hover{
    background-color:rgb(104, 104, 104);
    transition: 500ms;
}

.editor{
    margin: 15px;
}

.marked-enter-active {
    animation: marked-in .5s;
}
.marked-leave-active {
    animation: marked-in .5s reverse;
}
@keyframes marked-in {
    0% {
        transform: scale(0);
    }
    50% {
        transform: scale(1.2);
    }
    100% {
        transform: scale(1);
    }
}
</style>
