<template>
    <div>
        <div class="head-spacer"></div>
        <div class="title-content">
            Detail - {{mydetail.Title}}
        </div>
        <div :style="boximage">
            <div class="image-wrapper">
                click to fullscreen
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import {Vue, Component} from 'vue-property-decorator'
import axios from 'axios'

interface boxImage{
    width : string,
    height : string,
    minHeight : string,
    backgroundSize : string,
    backgroundPosition : string,
    "background-image" : string,
    margin : string
}

@Component({
    
})
export default class Detail extends Vue {
    id : string = ""
    mydetail : string = "aeiuo"
    path : string = ""
    boximage: boxImage = {
        width : "100vw",
        height : "auto",
        minHeight : "50vh",
        backgroundSize : "cover",
        backgroundPosition : "center",
        'background-image' : '',
        margin : "0 auto"
    }
    created() {
        this.id = this.$route.params.id
        axios.get('http://localhost:5000/api/images/'+this.id, {
            headers : {
                'Content-type' : 'application-json'
            }
        }).then(function (res){
            console.log(res.data)
            this.mydetail = res.data
            this.path = this.mydetail.Path
            Vue.set(this.boximage, 'backgroundImage', 'url('+require('@/server/'+this.path)+')')
        }.bind(this)).catch(function (err){
            console.log(err)
        })
    }
}
</script>

<style scoped>
.image-wrapper {
    position: fixed;
    z-index: 100;
    width: 100vw;
    height: 50vh;
    background: rgba(0, 0, 0, 0.5);
    color: white;
    font-weight: bold;
    font-size: 40px;
    display: flex;
    justify-content: center;
    align-items: center;
    text-decoration-line: underline;
}
</style>
