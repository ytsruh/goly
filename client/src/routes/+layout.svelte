<script>
  import "../app.css";
  import { feedback } from "../components/FeedbackStore";
  import Feedback from "../components/Feedback.svelte";
</script>

<div class="flex justify-between px-10 lg:px-20 p-2 bg-sky-500 text-white">
  <a href="/">Goly</a>
  <nav class="flex gap-10">
    <a href="/">Home</a>
    <a href="/create">Create</a>
  </nav>
</div>

{#if $feedback}
  <div class="flex justify-center w-full">
    <div class="mx-5 w-full md:w-3/4">
      <Feedback text={$feedback} />
    </div>
  </div>
{/if}

<slot />
