<template>
    <div class="cloth-gridbox">
        <cloth-box-item v-for="item in items" :key="item.id">
            {{item}}
        </cloth-box-item>
    </div>
</template>

<script lang="ts">
import {Vue, Component, Prop} from 'vue-property-decorator'
import boxItem from '@/views/ClothBoxItem.vue'

@Component({
    components: {
        'cloth-box-item' : boxItem
    }
})
export default class ClothBox extends Vue{
    @Prop() private path!: string;
    items : Array<string> = ["a", "b", "c", "d", "e", "f", "g"];
}
</script>

<style scoped>
.cloth-gridbox{
    display: grid;
    grid-gap: 10px;
    grid-template-columns: repeat(auto-fill, auto);
}

.cloth-item{
    background-color: white;
    border-radius: 10px;
    border-color: black;
    padding: 20px;
    height: 250px;
    max-width: 300px;
}
</style>
