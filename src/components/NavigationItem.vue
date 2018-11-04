<template lang="pug">
    .navcontent(@mouseover="openAccordion()")
        <slot></slot>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
@Component
export default class NavigationItem extends Vue{
    openAccordion() : void {
        this.$emit('open-accordion')
    }
}
</script>


<style scoped>
.navcontent{
    float: right;
    padding: 20px;
    cursor: pointer;
}
</style>
