from invoke import Collection, task


@task
def clean(ctx):
    """Remove all the untrack tmp files in"""
    ctx.run("git clean -f")


def strict_clean(ctx):
    """Strictly remove all the tmp files"""
    ctx.run("git clean -Xdf")


build_ns = Collection("build")
build_ns.add_task(clean)
