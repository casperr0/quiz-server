from invoke import task

from tasks.common import COMMON_MODULES_AS_PARAM, VENV_PREFIX


@task
def flake8(ctx):
    """Check style through flake8"""
    ctx.run(f"{VENV_PREFIX} flake8 --config=setup.cfg")


@task
def mypy(ctx):
    """Check style through mypy"""
    ctx.run(f"{VENV_PREFIX} mypy")


@task
def black_check(ctx):
    """Check style through black"""
    ctx.run(f"{VENV_PREFIX} black --check {COMMON_MODULES_AS_PARAM}")


@task
def isort_check(ctx):
    """Check style through isort"""
    ctx.run(f"{VENV_PREFIX} isort --atomic --check-only .")


@task
def commit_check(ctx):
    """Check commit message through commitizen"""
    result = ctx.run(f"{VENV_PREFIX} cz check --rev-range master..", warn=True)
    if result.exited == 3:  # NO_COMMIT_FOUND
        exit(0)
    else:
        exit(result.exited)


@task
def pylint(ctx):
    """Check style through pylint"""
    ctx.run(f"{VENV_PREFIX} pylint {COMMON_MODULES_AS_PARAM}")


# mypy, flake8, and pylint are not applicable for Django project
@task(pre=[black_check, isort_check, commit_check], default=True)
def run(ctx):
    """Check style through linters above"""
    pass


@task
def black(ctx):
    ctx.run(f"{VENV_PREFIX} black {COMMON_MODULES_AS_PARAM}")


@task
def isort(ctx):
    ctx.run(f"{VENV_PREFIX} isort --atomic .")


@task(pre=[black, isort])
def reformat(ctx):
    """Reformat python files throguh black and isort"""
    pass
