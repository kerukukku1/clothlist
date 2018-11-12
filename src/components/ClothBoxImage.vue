<template>
    <!-- <div class="cloth-box-image"> -->
    <div :style="boximage">
    </div>
</template>

<script lang="ts">
import {Vue, Component, Prop} from 'vue-property-decorator'

interface boxImage{
    width : string,
    height : string,
    minHeight : string,
    backgroundSize : string,
    backgroundPosition : string,
    "background-image" : string
}

@Component({})
export default class ClothBoxImage extends Vue{
    @Prop() imagePath! : string
    path : string = ""
    boximage: boxImage = {
        width : "100%",
        height : "200px",
        minHeight : "200px",
        backgroundSize : "cover",
        backgroundPosition : "center",
        'background-image' : ''
    }
    getImgUrl(data: string){
        return require(data)
    }
    mounted() {
        console.log(this.imagePath)
        this.path = this.imagePath
        Vue.set(this.boximage, 'backgroundImage', 'url('+require('@/server/'+this.path)+')')
    }
}
</script>

<style scoped>
.cloth-box-image{
    width: 100%;
    height: 200px;
    min-height: 100%;
    background-size: cover;
    background-position: center;
    background-image: url("../assets/users/mahito6/IMG_7696.jpg")
}

</style>
