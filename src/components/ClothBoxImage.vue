<template>
    <!-- <div class="cloth-box-image"> -->
    <div :style="boximage">
    </div>
</template>

<script lang="ts">
import {Vue, Component} from 'vue-property-decorator'

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
    boximage: boxImage = {
        width : "100%",
        height : "200px",
        minHeight : "200px",
        backgroundSize : "cover",
        backgroundPosition : "center",
        'background-image' : ''
    }
    created() {
        Vue.set(this.boximage, 'backgroundImage', 'url('+require("../assets/users/mahito6/IMG_7696.jpg")+')')
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
