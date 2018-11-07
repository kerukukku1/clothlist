<template>
    <div>
        <canvas width="0" height="0" ref="my-canvas"></canvas>
    </div>    
</template>

<script lang="ts">
import {Vue, Component} from 'vue-property-decorator'
@Component({

})
export default class PreviewCanvas extends Vue {
    onPreview(file : Blob){
        console.log("preview!")
        var blobUrl = URL.createObjectURL( file )
        var image : HTMLImageElement = new Image
        image.src = blobUrl;
        image.onload = function(){
            var canvas : HTMLCanvasElement = this.$refs['my-canvas'];
            console.log(canvas)
            console.log(image)
            console.log(image.width)
            var context : CanvasRenderingContext2D = canvas.getContext('2d')
            canvas.width = image.width;
            canvas.height = image.height;
            context.drawImage(image, 0, 0, canvas.width, canvas.height);
        }.bind(this)
    }
}
</script>

<style scoped>

</style>
