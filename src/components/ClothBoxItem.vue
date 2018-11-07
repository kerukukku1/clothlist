<template>
    <div class="cloth-item">
        <cloth-box-image></cloth-box-image>
        <hr>
        <slot></slot>
    </div>
</template>

<script lang="ts">
import {Vue, Component} from 'vue-property-decorator'
import boxImage from '@/components/ClothBoxImage.vue'

@Component({
    components:{
        'cloth-box-image' : boxImage
    }
})
export default class ClothBoxItem extends Vue{

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
