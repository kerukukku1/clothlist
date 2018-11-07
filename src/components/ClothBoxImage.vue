<template>
    <div class="cloth-box-image">
    </div>
</template>

<script lang="ts">
import {Vue, Component} from 'vue-property-decorator'

@Component({})
export default class ClothBoxImage extends Vue{

}
</script>

<style scoped>
.cloth-box-image{
    width: 100%;
    height: 200px;
    min-height: 100%;
    background-size: cover;
    background-position: center;
    background-image: url("../assets/users/mahito6/IMG_7696.jpg");
}

</style>
