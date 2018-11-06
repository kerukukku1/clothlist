<template>
    <div class="cloth-item">
        <slot></slot>
    </div>
</template>

<script lang="ts">
import {Vue, Component} from 'vue-property-decorator'

@Component({})
export default class ClothBoxItem extends Vue{

}
</script>

<style scoped>
.cloth-item{
    background-color: white;
    border-radius: 10px;
    border-color: black;
    padding: 20px;
    height: 200px;
}
</style>
