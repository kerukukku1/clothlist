<template>
    <div class="home">
        <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
        <div class="head-spacer"></div>
        <div class="title-content">
            Upload Form
        </div>
        <!-- <form action="http://localhost:5000/images/post" method="post" enctype="multipart/form-data" onsubmit="return false;"> -->
            <input type="file" name="imagefile" mulitple="multiple" accept="image/*" @change="onDrop">
            <input type="text" name="title" v-model="title">
            <input type="submit" name="submit" value="アップロード開始" @click="onSubmit">
        <!-- </form> -->
        <preview-canvas ref="preview"></preview-canvas>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import previewCanvas from '@/components/PreviewCanvas.vue'
import axios from 'axios';

@Component({
    components: {
        'preview-canvas' : previewCanvas
    },
})
export default class Upload extends Vue {
    title : string = ""
    file : Blob = new Blob();

    getBase64(file : Blob) {
        var reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = function () {
            console.log(reader.result);
        };
        reader.onerror = function (error) {
            console.log('Error: ', error);
        };
    }

    onSubmit(){
        console.log("submit!")
        let postData = new FormData()
        postData.append('title', this.title)
        postData.append('imagefile', this.file)
        const options = {
            headers : {
                'content-type': 'multipart/form-data',
            }
        }
        // axios.post('http://192.168.11.8:5000/images/post', postData, options)
        axios.post('http://localhost:5000/images/post', postData, options)
        .then(function (res){
            console.log(res)
        })
        .catch(function (err){
            console.log(err)
        })
        // const options = {
        //     method: 'POST',
        //     headers: { 
        //         'Content-type' : "application/json"
        //     },
        //     data: this.file,
        //     url : 'http://localhost:5000/images/post'
        // };
        // let params = new URLSearchParams();
        // params.append('body', JSON.stringify(this.file));
        // axios.post('http://localhost:5000/images/post', params).then(response => {
        //     console.log(response)
        // })
        // .catch(e => {
        //     console.log(e)
        // })
    }

    onDrop(event : Event){
        // var blobUrl = URL.createObjectURL( event.target.files[0] ) ;
        // var xhr = new XMLHttpRequest();
        // xhr.onload = function() {
        //     var result = xhr.response;
        //     this.getBase64(result);
        // }.bind(this);
        // xhr.responseType = "blob";
        // xhr.open("GET", blobUrl);
        // xhr.send();
        this.file = event.target.files[0]
        this.$refs.preview.onPreview(this.file)
        console.log("change detect")
    }
    created(){
        console.log("upload created")
    }
}
</script>

<style scoped>
</style>
