# frozen_string_literal: true
# typed: strict
Spring.watch(
  ".ruby-version",
  ".rbenv-vars",
  "tmp/restart.txt",
  "tmp/caching-dev.txt",

)

# https://github.com/TalentBox/sequel-rails/commit/e3e35209ee45918ce1f61daa6ef7850c881fac24
Spring.after_fork do
  Sequel::DATABASES.each(&:disconnect)
end
