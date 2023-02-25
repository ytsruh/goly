<script lang="ts">
  export let text: string;
  import { feedback } from "./FeedbackStore";
  import FaWindowClose from "svelte-icons/fa/FaWindowClose.svelte";
  function dismiss() {
    feedback.set("");
  }
  setTimeout(() => {
    feedback.set("");
  }, 3000);
</script>

<div class="flex justify-between bg-sky-900 my-2 py-2 px-6 rounded-md">
  <h2 class="text-xl text-white text-center ">{text}</h2>
  <div class="cursor-pointer text-white pt-1 h-6 w-6" on:click={dismiss} on:keypress={dismiss}>
    <FaWindowClose />
  </div>
</div>
