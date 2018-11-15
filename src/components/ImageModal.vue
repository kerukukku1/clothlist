<template>
    <transition name="modal" appear>
        <div class="modal-overlay" @click.self="$emit('close')">
            <div class="modal-window"  @click.self="$emit('close')">
                <div class="modal-content"  @click.self="$emit('close')">
                    <slot></slot>
                </div>
            </div>
        </div>
    </transition>    
</template>

<script lang="ts">
import {Vue, Component} from 'vue-property-decorator'
@Component({

})
export default class ImageModal extends Vue{

}
</script>

<style scoped>
.modal-overlay {
    display: flex;
    align-items: center;
    justify-content: center;
    position: fixed;
    z-index: 10000;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background: rgba(0, 0, 0, 0.5);
}

.modal-window{
    width: auto;
    height: auto;
}

.modal-enter-active, .modal-leave-active {
    transition: opacity 0.3s;
}

.modal-leave-active {
    transition: opacity 0.3s;
}

.modal-enter, .modal-leave-to {
    opacity : 0;
}


</style>
