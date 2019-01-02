#### Why?
The reasons which I choose Ruby for the automation testing:
* Ruby is better language for programming
* Easier to approach the automation test for the new guys

#### What?
Using Ruby programming language and the relate extensions for automation testing

#### How?
The way these component work together:

Ruby env <==> Bundler <==> RSpec <==> Capybara <==> Selenium <==> Browser

* [Bundler](http://bundler.io/) provides a consistent environment for Ruby projects.
* [RSpec](http://rspec.info) is a testing tool for Ruby, created for behavior-driven development (BDD), look like PHP_UNIT in PHP.
* [Capybara](https://github.com/teamcapybara/capybara) is an acceptance test framework for web applications.
* Selenium is a suite of tools to automate web browsers across many platforms.

#### Install
* [Homebrew](https://brew.sh/)
* Xcode
* [Ruby Version Manager (RVM)](https://rvm.io/)

```
\curl -L https://get.rvm.io | bash -s stable
source $HOME/.rvm/scripts/rvm
#Run rvm requirements to install additional dependencies:   
rvm requirements
```

* Ruby

```
rvm install 2.4.0
rvm use 2.4.0 --default
```

* Install gem [bundler](http://bundler.io/) (provides a consistent environment for Ruby projects by tracking and installing the exact gems and versions that are needed)

```
gem install bundler
```

* Setup browser drivers

```
#For Gecko driver :(using for Firefox browser)
brew install geckodriver
#For Chrome driver:
brew install chromedriver
#In addition, we must Enable Automation mode on Safari browser
#Open to Safari > Preferences
#Check on "Show develop menu in menu bar"
#Then open menu develop and Check on “Allow Remote Automation”
```

#### Cheat sheet

**Navigating**
```ruby
visit '[REQUEST_URI]'
```

**Scoping**
```ruby
within('head', visible: false) do
    expect(page).to have_xpath('//meta[@name="description"][@content="' + META_DES + '"]', :visible => false)
    expect(page).to have_xpath('//meta[@name="keywords"][@content="' + META_KEYWORDS + '"]', :visible => false)
end
```
```ruby
within(:xpath, '//div[@class="article-cta"]') do
  click 'Sign up'
end
```
```ruby
within(:xpath, '//nav')

within_fieldset('Employee') do
  fill_in 'Name', with: 'Jimmy'
end

within_table('Employee') do
  fill_in 'Name', with: 'Jimmy'
end

within_frame

find('div#article-cta').click 'Sign up'
```

**Clicking links and buttons [Capybara::Node::Actions](http://rubydoc.info/github/teamcapybara/capybara/master/Capybara/Node/Actions)**

```ruby
click_link('id-of-link')
click_link('Link Text')
click_button('Save')
click_on('Link Text') # clicks on either links or buttons
click_on('Button Value')
```

**Interacting with forms [Capybara::Node::Actions](http://rubydoc.info/github/teamcapybara/capybara/master/Capybara/Node/Actions)**
```ruby
fill_in('First Name', with: 'John')
fill_in('Password', with: 'Seekrit')
fill_in('Description', with: 'Really Long Text...')
choose('A Radio Button')
check('A Checkbox')
uncheck('A Checkbox')
attach_file('Image', '/path/to/image.jpg')
select('Option', from: 'Select Box')
```

**Querying [Capybara::Node::Matchers](http://rubydoc.info/github/teamcapybara/capybara/master/Capybara/Node/Matchers)**

```ruby
page.has_selector?('table tr')
page.has_selector?(:xpath, './/table/tr')

page.has_xpath?('.//table/tr')
page.has_css?('table tr.foo')
page.has_content?('foo')
```

**Note:** The negative forms like `has_no_selector?` are different from `not
has_selector?`. Read the section on asynchronous JavaScript for an explanation.

You can use these with RSpec's magic matchers:

```ruby
expect(page).to have_selector('table tr')
expect(page).to have_selector(:xpath, './/table/tr')

expect(page).to have_xpath('.//table/tr')
expect(page).to have_css('table tr.foo')
expect(page).to have_content('foo')
```

**Finding [Capybara::Node::Finders](http://rubydoc.info/github/teamcapybara/capybara/master/Capybara/Node/Finders)**

```ruby
find_field('First Name').value
find_field(id: 'my_field').value
find_link('Hello', :visible => :all).visible?
find_link(class: ['some_class', 'some_other_class'], :visible => :all).visible?

find_button('Send').click
find_button(value: '1234').click

find(:xpath, ".//table/tr").click
find("#overlay").find("h1").click
all('a').each { |a| a[:href] }
```

**Working with windows**
```ruby
facebook_window = window_opened_by do
  click_button 'Like'
end
within_window facebook_window do
  find('#login_email').set('a@example.com')
  find('#login_password').set('qwerty')
  click_button 'Submit'
end
```

**Assertions**
```ruby
page.has_button?('Save')
expect(page).to have_no_button('Save')
expect(page).to have_button('Save')
expect(page).to have_button('#submit')
expect(page).to have_button('//[@id="submit"]')
```

**Common options**
```ruby
text: 'welcome'
text: /Hello/
visible: true
count: 4
between: 2..5
minimum: 2
maximum: 5
wait: 10
```

```ruby
expect(page).to have_link('Logout', href: logout_path)
expect(page).to have_select('Language',
                            selected: 'German'
                            options: ['Engish', 'German']
                            with_options: ['Engish', 'German'])
```

**Page**
```ruby
page
  .all('h3')
  .body
  .html
  .source
  .current_host
  .current_path
  .current_url
  
page.status_code == 200
page.response_headers  
```

**AJAX**
```ruby
using_wait_time 10 do
  ...
end
```

**Modals**
```ruby
accept_alert do
  click_link('Show Alert')
end
```

```ruby
dismiss_confirm do
  click_link('Show Confirm')
end
```

```ruby
accept_prompt(with: 'Linus Torvalds') do
  click_link('Show Prompt About Linux')
end
```

```ruby
message = accept_prompt(with: 'Linus Torvalds') do
  click_link('Show Prompt About Linux')
end
expect(message).to eq('Who is the chief architect of Linux?')
```

