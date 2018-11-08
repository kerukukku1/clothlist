<template>
    <div class="home">
        <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
        <div class="head-spacer"></div>
        <div class="title-content">
            Upload Form
        </div>
        <input type="file" mulitple="multiple" accept="image/*" @change="onDrop">
        <input type="submit" name="submit" value="アップロード開始" @click="onSubmit">
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
        axios.post('http://localhost:5000/images/post', {
            headers : {
                'Content-type' : "application/json"
            }
        }).then(response => {
            console.log(response)
        })
        .catch(e => {
            console.log(e)
        })
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
