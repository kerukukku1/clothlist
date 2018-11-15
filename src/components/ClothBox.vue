<template>
    <div class="cloth-gridbox">
        <cloth-box-item v-for="item in images" :key="item.id" :imageData="item">
            <a :href="'/gallery/'+item.ID"> {{item.Title}} </a>
        </cloth-box-item>
    </div>
</template>

<script lang="ts">
import {Vue, Component, Prop, Watch} from 'vue-property-decorator'
import boxItem from '@/components/ClothBoxItem.vue'
import axios from 'axios'

interface myImage {
    ID : string,
    Title : string,
    Path : string
}

@Component({
    components: {
        'cloth-box-item' : boxItem
    }
})
export default class ClothBox extends Vue{
    images : Array<myImage> = []

    mounted() {
        axios.get('http://localhost:5000/api/images', {
            headers : {
                'Content-type' : 'application/json'
            }
        }).then(function (res){
            this.images = res.data
            for(let index in this.images){
                this.images[index].Path = this.images[index].Path
            }
        }.bind(this)).catch(function (err){
            console.log(err)
        })
    }
}
</script>

<style scoped>
.cloth-gridbox{
    display: grid;
    grid-gap: 10px;
    grid-template-columns: repeat(auto-fill, auto);
}

a {
    color: black;
    font-weight: bold;
}

</style>
