<script>
  import "../app.css";
</script>

<div class="flex justify-between px-10 lg:px-20 p-2 bg-sky-500 text-white">
  <a href="/">Goly</a>
  <nav class="flex gap-10">
    <a href="/">Home</a>
    <a href="/create">Create</a>
  </nav>
</div>

<slot />
