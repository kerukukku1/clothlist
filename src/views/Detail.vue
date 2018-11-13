<template>
    <div>
        {{mydetail}}
    </div>
</template>

<script lang="ts">
import {Vue, Component} from 'vue-property-decorator'
import axios from 'axios'

@Component({
    
})
export default class Detail extends Vue {
    id : string = ""
    mydetail : string = "aeiuo"
    created() {
        this.id = this.$route.params.id
        axios.get('http://localhost:5000/api/images/'+this.id, {
            headers : {
                'Content-type' : 'application-json'
            }
        }).then(function (res){
            console.log(res.data)
            this.mydetail = res.data
        }.bind(this)).catch(function (err){
            console.log(err)
        })
    }
}
</script>

