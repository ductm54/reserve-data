FROM ruby:2.5

RUN apt-get update && apt-get install -y git coreutils make g++ nodejs
RUN git clone https://github.com/lord/slate /slate
COPY ./index.html.md /slate/source/
WORKDIR /slate
RUN bundle install

CMD exec bundle exec middleman build --clean 