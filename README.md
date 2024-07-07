# Persian Gender Detection

A Go package for detecting gender by Persian and Arabic names (with more than 23K names).

## Installation

`go get -u github.com/armanyazdi/persian-gender-detection-go`

## Usage

Let's take a look at what an example test case would look like using `persian-gender-detection`.

### Detect Gender:

```go
import
genderdetector "github.com/armanyazdi/persian-gender-detection-go"
)
```

```go
genderdetector.GetGender("  عــــلی  ")    // MALE
genderdetector.GetGender("محــ🌚ــمد")     // MALE
genderdetector.GetGender("بیــ🥲ــتا")     // FEMALE
genderdetector.GetGender("۱۲۳مهناز۴۵۶")    // FEMALE
genderdetector.GetGender("مهدي-1980")      // MALE
genderdetector.GetGender("بابك آذر مهر")   // MALE
genderdetector.GetGender("بی بی سلطان")    // FEMALE
genderdetector.GetGender("شیدا علیزاده")   // FEMALE
genderdetector.GetGender("ممد رضا")        // MALE
genderdetector.GetGender("پانته‌آ عَبّاسی")   // FEMALE
genderdetector.GetGender("دکتر ندا حسینی") // FEMALE
genderdetector.GetGender("سید امیر موسوی") // MALE
genderdetector.GetGender("جناب آقای X")    // MALE
genderdetector.GetGender("سرکار خانم Y")   // FEMALE
```

## Issues

Feel free to submit issues and enhancement requests.

## Contributing

Feel free to contribute names database with your kindly pull requests.

## TODO

- [ ] Add Finglish support

## License

`persian-gender-detection` is available under the [MIT license](https://github.com/armanyazdi/persian-gender-detection-go/blob/master/LICENSE).