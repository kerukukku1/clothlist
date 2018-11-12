<template>
    <div class="cloth-item">
        <cloth-box-image :imagePath="imageData.Path"></cloth-box-image>
        <hr>
        <slot></slot>
    </div>
</template>

<script lang="ts">
import {Vue, Component, Prop} from 'vue-property-decorator'
import boxImage from '@/components/ClothBoxImage.vue'

interface myImage {
    ID : string,
    Title : string,
    Path : string
}

@Component({
    components:{
        'cloth-box-image' : boxImage
    }
})
export default class ClothBoxItem extends Vue{
    @Prop() imageData!: myImage;
    mounted() {
    }
}
</script>

<style scoped>
.cloth-item{
    background-color: azure;
    border-radius: 10px;
    border-color: black;
    padding: 10px;
    height: auto;
}
</style>
