<template>
    <div>
        <canvas width="0" height="0" ref="my-canvas" v-bind:style="mystyle"></canvas>
    </div>    
</template>

<script lang="ts">
import {Vue, Component} from 'vue-property-decorator'
const THUMBNAIL_WIDTH = 600;
const THUMBNAIL_HEIGHT = 600;
function getResizeImageSize(image : HTMLImageElement){
    if(image.width > image.height){
        // 横長の画像は横のサイズを指定値にあわせる
        var ratio = image.height/image.width;
        var width = THUMBNAIL_WIDTH;
        var height = THUMBNAIL_WIDTH * ratio;
    } else {
        // 縦長の画像は縦のサイズを指定値にあわせる
        var ratio = image.width/image.height;
        var width = THUMBNAIL_HEIGHT * ratio;
        var height = THUMBNAIL_HEIGHT;
    }
    return {width, height}
}

interface styleObject{
    width : String,
    height : String
}

@Component({

})
export default class PreviewCanvas extends Vue {
    mystyle : styleObject = {
        width : "0px",
        height : "0px"
    }
    onPreview(file : Blob){
        console.log("preview!")
        var blobUrl = URL.createObjectURL( file )
        var image : HTMLImageElement = new Image
        image.src = blobUrl;
        image.onload = function(){
            var canvas : HTMLCanvasElement = this.$refs['my-canvas'];
            var context : CanvasRenderingContext2D = canvas.getContext('2d')
            canvas.width = image.width;
            canvas.height = image.height;
            var size = getResizeImageSize(image)
            canvas.width = size.width;
            canvas.height = size.height;
            Vue.set(this.mystyle, 'width', (canvas.width/2)+'px')
            Vue.set(this.mystyle, 'height', (canvas.height/2)+'px')
            context.drawImage(image,0,0,image.width,image.height,0,0,canvas.width,canvas.height);
            var base64 : String  = canvas.toDataURL("image/jpeg")
            console.log(base64)
        }.bind(this)
    }
}
</script>

<style scoped>

</style>
