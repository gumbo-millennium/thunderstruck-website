<template>
  <div class="font-archivo relative text-3xl uppercase font-black">
    <h1 class="absolute left-[-.2rem] top-[-.1rem] z-10">
      <slot />
    </h1>
    <h1
      aria-hidden="true"
      class="text-white stroke-black drop-shadow-[0_0.5px_0.5px_rgba(0,0,0,1)] z-0"
    >
      <slot />
    </h1>
  </div>
</template>
