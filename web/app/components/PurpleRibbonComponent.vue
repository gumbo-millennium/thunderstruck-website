<template>
    <div class="rotate-40 absolute text-white font-black font-kodchasan text-2xl overflow-hidden right-[-11rem] z-10">
      <div class="relative">
        <div class="bg-purple flex items-center overflow-x-hidden h-10 animate-horizontal-scrolling whitespace-nowrap"
        >
          <span><slot /></span>
          <span><slot /></span>
        </div>
        <div class="absolute top-0 bg-purple flex items-center overflow-x-hidden h-10 animate-horizontal-scrolling-2 whitespace-nowrap">
          <span><slot /></span>
          <span><slot /></span>
        </div>
        <div class="absolute top-0 bg-purple flex items-center overflow-x-hidden h-10 animate-horizontal-scrolling-3 whitespace-nowrap">
          <span><slot /></span>
          <span><slot /></span>
        </div>
      </div>
    </div>
</template>

<script setup lang="ts">
</script>
