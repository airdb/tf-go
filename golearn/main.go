package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/linear_models"
)

func main() {
	fmt.Println("xx")

	trainData, _ := base.ParseCSVToInstances("./datasets/exams.csv", true)
	testData, _ := base.ParseCSVToInstances("./datasets/exam.csv", true)

	lr := linear_models.NewLinearRegression()
	lr.Fit(trainData)

	a, _ := lr.Predict(testData)
	fmt.Println(a)

	/*
	return
	testData, err := base.ParseCSVToInstances("./datasets/exams.csv", true)
	fmt.Println(testData, err)

	lr := linear_models.NewLinearRegression()
	trainingDatum, err := base.ParseCSVToInstances("./datasets/exam.csv", true)


	lr.Fit(trainingDatum)

	a, err := lr.Predict(testData)
	fmt.Println(a, err)
	Convey("Doing a  linear regression", func() {

		Convey("With no training data", func() {
			Convey("Predicting", func() {
				testData, err := base.ParseCSVToInstances("../examples/datasets/exams.csv", true)
				So(err, ShouldBeNil)

				_, err = lr.Predict(testData)

				Convey("Should result in a NoTrainingDataError", func() {
					So(err, ShouldEqual, linear_models.NoTrainingDataError)
				})

			})
		})

		Convey("With not enough training data", func() {
			trainingDatum, err := base.ParseCSVToInstances("../examples/datasets/exam.csv", true)
			So(err, ShouldBeNil)

			Convey("Fitting", func() {
				err = lr.Fit(trainingDatum)

				Convey("Should result in a NotEnoughDataError", func() {
					So(err, ShouldEqual, linear_models.NotEnoughDataError)
				})
			})
		})

		Convey("With sufficient training data", func() {
			instances, err := base.ParseCSVToInstances("../examples/datasets/exams.csv", true)
			So(err, ShouldBeNil)
			trainData, testData := base.InstancesTrainTestSplit(instances, 0.1)

			Convey("Fitting and Predicting", func() {
				err := lr.Fit(trainData)
				So(err, ShouldBeNil)

				predictions, err := lr.Predict(testData)
				So(err, ShouldBeNil)

				Convey("It makes reasonable predictions", func() {
					_, rows := predictions.Size()

					for i := 0; i < rows; i++ {
						actualValue, _ := strconv.ParseFloat(base.GetClass(testData, i), 64)
						expectedValue, _ := strconv.ParseFloat(base.GetClass(predictions, i), 64)

						So(actualValue, ShouldAlmostEqual, expectedValue, actualValue*0.05)
					}
				})
			})
		})
	})
	 */
}

/*
func BenchmarkLinearRegressionOneRow(b *testing.B) {
	// Omits error handling in favor of brevity
	trainData, _ := base.ParseCSVToInstances("../examples/datasets/exams.csv", true)
	testData, _ := base.ParseCSVToInstances("../examples/datasets/exam.csv", true)
	lr := NewLinearRegression()
	lr.Fit(trainData)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		lr.Predict(testData)
	}
}
 */