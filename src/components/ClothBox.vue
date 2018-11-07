<template>
    <div class="cloth-gridbox">
        <cloth-box-item v-for="item in items" :key="item.id">
            {{item}} => Show Details
        </cloth-box-item>
    </div>
</template>

<script lang="ts">
import {Vue, Component, Prop} from 'vue-property-decorator'
import boxItem from '@/components/ClothBoxItem.vue'

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

</style>
