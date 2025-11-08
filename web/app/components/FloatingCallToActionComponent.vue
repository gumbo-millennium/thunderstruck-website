<template>
  <div class="fixed bottom-0 right-0 m-8 bg-green-secondary p-8 animate-bounce">
    <CallToActionComponent
      :href="props.url"
    >
      <slot />
    </CallToActionComponent>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  url: string,
}>();
</script>
