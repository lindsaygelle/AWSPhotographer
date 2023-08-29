Contributing Code
-----------------
A good pull request:

-  Is clear.
-  Works across all supported versions.
-  Follows the existing style of the code base.
-  Has comments included as needed.

-  A test case that demonstrates the previous flaw that now passes with
   the included patch, or demonstrates the newly added feature.
-  If it adds/changes a public API, it must also include documentation
   for those changes.
-  Must be appropriately [licensed](./LICENSE).

Reporting An Issue/Feature
--------------------------
First, check to see if there's an existing issue/pull request for the
bug/feature. All issues are at
https://github.com/lindsaygelle/AWSPhotographer/issues and pull reqs are at
https://github.com/lindsaygelle/AWSPhotographer/pulls.

If there isn't an existing issue there, please file an issue. The
ideal report includes:

-  A description of the problem/suggestion.
-  How to recreate the bug.
-  If possible, create a pull request with a (failing) test case
   demonstrating what's wrong. This makes the process for fixing bugs
   quicker & gets issues resolved sooner.

Codestyle
---------
This project uses pre-commit to enforce codstyle requirements.

To validate your PR prior to publishing, you can use the following
`installation guide <https://pre-commit.com/#install>`__ to setup pre-commit.

If you don't want to use the git commit hook, you can run the below command
to automatically perform the codestyle validation:

.. code-block:: bash

    $ pre-commit run --all-files

This will automatically perform simple updates (such as white space clean up)
and provide a list of any formatting checks. After these are addressed,
you can commit the changes prior to publishing the PR.
