package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"kanmaa-backend/graph/generated"
	"kanmaa-backend/graph/model"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func (r *mutationResolver) AddFile(ctx context.Context, file graphql.Upload, fileType model.FileTypeEnum, requiredByID string, uploader string, fileURI string, schoolEmail string) (*model.StaticFile, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddUser(ctx context.Context, email string, password string, salute model.SaluteEnum, name string, iDno string, nHIFno string, cell string, gender model.GenderEnum, school string, addrPoBox string, addrPostalCode string, addrDistrict string, addrCounty string, addrNationality string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddStudent(ctx context.Context, name string, nemisNo string, nhifStatus bool, bCertNo string, nickName string, dob time.Time, hoby1 string, hoby2 string, hoby3 string, roleModel1 string, roleModel2 string, roleModel3 string, career1 string, career2 string, career3 string, motto string, bloodGroup string, bestFriend string, language1 string, language2 string, language3 string, game string, subject string, book string, quote string, animal string, device string, grade model.GradeEnum, school string, parent string, fees float64, sharemeal bool, greetstranger bool, firstaider bool) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddSchool(ctx context.Context, name string, motto string, phone string, email string, website string, addrPoBox string, addrPostalCode string, addrTown string, addrCounty string, addrCountry string) (*model.School, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddDepartment(ctx context.Context, name model.DeptEnum, hod string, school string) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddSubject(ctx context.Context, name model.SubjectEnum, department model.DeptEnum, schoolEmail string) (*model.Subject, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddGrade(ctx context.Context, name model.GradeEnum, stream string, classTeacher string, motto string, schoolEmail string) (*model.Grade, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddLesson(ctx context.Context, name string, date string, timeDate model.PeriodEnum, tutorEmail string, grade model.GradeEnum, subject model.SubjectEnum, start string, stop string, duration float64, done bool, remarks string, school string) (*model.Lesson, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddTopic(ctx context.Context, lesson string, strand string, subStrand string, outComes string, experiences string, inquiries string, competencies string, lifeSkills string, esddrr string, values string, otherAreas string, communityActivities string, nonFormalActivity string, assessment string, schoolEmail string) (*model.Lesson, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddBook(ctx context.Context, name string, isbn string, authors string, revision string, school string) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LendOut(ctx context.Context, bookIsbn string, receivedOn string, dueOn string, student string, librarian string, school string) (*model.Lend, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ReceiveBook(ctx context.Context, lendID string, condition model.ConditionEnum) (*model.Lend, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddFee(ctx context.Context, grade model.GradeEnum, year string, amount float64, semester string, school string) (*model.Fee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) NewPayment(ctx context.Context, paytype model.PaymentTypeEnum, payFor model.PaymentForEnum, receiver string, amount float64, date string, account model.AccEnum, payeemail string, school string) (*model.Payment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddExam(ctx context.Context, name string, date string, grade model.GradeEnum, school string) (*model.Exam, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddScore(ctx context.Context, score float64, studentNemis string, subject model.SubjectEnum, examID string, schoolEmail string) (*model.Score, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddTimeTable(ctx context.Context, day string, lesson1 string, lesson2 string, lesson3 string, lesson4 string, lesson5 string, lesson6 string, lesson7 string, lesson8 string, grade model.GradeEnum, schoolEmail string) (*model.TimeTable, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) NewVisitor(ctx context.Context, salute string, name string, iDno string, cell string, email string, visiting string, purpose string, vehicle string, belongings string, visitingFreq float64, authorizedBy string, schoolEmail string) (*model.Visitor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) NewDiary(ctx context.Context, tutorRemarks string, student string, schoolRequest string, tutorMail string, schoolEmail string) (*model.Diary, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ViewDiary(ctx context.Context, guardianEmail string, diaryID string, schoolEmail string) (*model.Diary, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddProduct(ctx context.Context, name string, cost float64, quantity float64, supplier string, category model.ProductCategoryEnum, receivedOn string, receivedBy string, allowedMinimum float64, schoolEmail string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DispenseProduct(ctx context.Context, name string, category string, quantity float64, dispenseTo string, dispensedBy string, schoolEmail string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) NewCulture(ctx context.Context, title string, instructions string, story string, quiz string, lessons string, culture string, grade model.GradeEnum, schoolEmail string) (*model.Culture, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) NewEssene(ctx context.Context, userMail string, schoolEmail string) (*model.Essene, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StudentCheckin(ctx context.Context, nemisID string, schoolEmail string) (*model.StudentinStudentout, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StudentCheckout(ctx context.Context, sisoID string) (*model.StudentinStudentout, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StaffCheckin(ctx context.Context, workerEmail string, authorizedBy string, schoolEmail string) (*model.CheckinCheckout, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StaffCheckout(ctx context.Context, cicoID string) (*model.CheckinCheckout, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SendAnnouncement(ctx context.Context, topic string, by string, message string, schoolEmail string) (*model.Announcement, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ReadAnnouncement(ctx context.Context, announceID string, user string, read bool, schoolEmail string) (*model.Announcement, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SendMessage(ctx context.Context, topic string, text string, user string, thread string, school string) (*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) VisitorCheckout(ctx context.Context, visitorID string) (*model.Visitor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddRole(ctx context.Context, userEmail string, roleName model.RolesEnum) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RestoreRoles(ctx context.Context, userEmail string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PromoteStudent(ctx context.Context, grade model.GradeEnum, studentNemisID string, schoolEmail string) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ConfirmPayment(ctx context.Context, payID string, email string) (*model.Payment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ConfirmLessonAttended(ctx context.Context, attendanceID string, studentGif graphql.Upload) (*model.Attendance, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AttendLesson(ctx context.Context, nemisID string, lessonID string, school string) (*model.Attendance, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddStudentGuardian(ctx context.Context, studentNemis string, guardianEmail string, schoolEmail string) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelAnnouncement(ctx context.Context, announcementID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelAnnounceUserConn(ctx context.Context, connID string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelAttendance(ctx context.Context, attendanceID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelBook(ctx context.Context, bookID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelCico(ctx context.Context, cicoID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelSiso(ctx context.Context, sisoID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelParentingCulture(ctx context.Context, parentingCultureID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelDepartment(ctx context.Context, deptID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelDiary(ctx context.Context, diaryID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelUserDiaryConn(ctx context.Context, connID string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelDispensedProduct(ctx context.Context, disProductID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelProduct(ctx context.Context, prodID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelEssene(ctx context.Context, esseneID string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelExam(ctx context.Context, examID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelFee(ctx context.Context, feeID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelGrade(ctx context.Context, gradeID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelLend(ctx context.Context, lendID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelLesson(ctx context.Context, lssnID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelMessage(ctx context.Context, msgID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelPayment(ctx context.Context, payID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelSchool(ctx context.Context, schoolID string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelScore(ctx context.Context, scoreID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelStudent(ctx context.Context, nemisID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelSubject(ctx context.Context, subjectID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelThread(ctx context.Context, threadID string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelTimeTable(ctx context.Context, ttableID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelUser(ctx context.Context, userEmail string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelVisitor(ctx context.Context, visitorID string, deletedByName string, deletedByMail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditAnnouncement(ctx context.Context, announcementID string, topic *string, announcerEmail *string, message *string) (*model.Announcement, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditBook(ctx context.Context, bookID string, name *string, isbn *string, authors *string, revision *string) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditCulture(ctx context.Context, parentingCultureID string, storyTitle *string, instructions *string, story *string, quiz *string, lessonsLearnt *string, cultureTaught *string, gradeToReceive *model.GradeEnum) (*model.Culture, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditDepartment(ctx context.Context, deptID string, name *model.DeptEnum, hodEmail *string) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditDiary(ctx context.Context, diaryID string, remarks *string, studentNemis *string, request *string, tutorEmail *string) (*model.Diary, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditDispensedProduct(ctx context.Context, disProductID string, productName *string, quantity *float64, category *model.ProductCategoryEnum, dispensedTo *string, dispensedBy *string) (*model.DispensedProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditProduct(ctx context.Context, prodID string, name *string, cost *float64, quantity *float64, supplier *string, category *model.ProductCategoryEnum, receivedOn *string, receivedBy *string, minimumQty *float64) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditExam(ctx context.Context, examID string, name *string, date *time.Time, gradeName *model.GradeEnum) (*model.Exam, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditFee(ctx context.Context, feeID string, grade *model.GradeEnum, year *string, amount *float64, semester *model.SenesterEnum) (*model.Fee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditGrade(ctx context.Context, gradeID string, name *model.GradeEnum, stream *string, classTeacher *string, gradeMotto *string) (*model.Grade, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditNewsletter(ctx context.Context, newsletterID string, mainTopic *string, topicOne *string, textOne *string, topicTwo *string, textTwo *string, topicThree *string, textThree *string, topicFour *string, textFour *string, topicFive *string, textFive *string, rubberStampURL *string) (*model.Newsletter, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditLesson(ctx context.Context, lssnID string, name *string, date *time.Time, timedate *model.PeriodEnum, grade *model.GradeEnum, subject *model.SubjectEnum, start *time.Time, stop *time.Time, duration *float64, done *bool, remarks *string, tutorEmail *string) (*model.Lesson, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditMessage(ctx context.Context, msgID string, topic *string, text *string, userEmail *string, threadID *string) (*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditPayment(ctx context.Context, payID string, paymentType *model.PaymentTypeEnum, payingFor *model.PaymentForEnum, receiverEmailNemisid *string, amount *float64, date *time.Time, account *model.AccEnum, payeeEmail *string) (*model.Payment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditSchool(ctx context.Context, schoolID string, name *string, motto *string, phone *string, email *string, website *string, poBox *string, poCode *string, district *string, county *string, country *string) (*model.School, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditScore(ctx context.Context, scoreID string, score *float64, studentNemisID *string, subject *model.SubjectEnum, examID *string) (*model.Score, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditStudent(ctx context.Context, stdID string, name *string, nemisNo *string, nhif *string, bCertNo *string, nickName *string, dateOfBirth *time.Time, hoby1 *string, hoby2 *string, hoby3 *string, roleModel1 *string, roleModel2 *string, roleModel3 *string, career1 *string, career2 *string, career3 *string, personalMotto *string, bloodGroup *string, bestFriend *string, language1 *string, language2 *string, language3 *string, quickAnswer1 *string, quickAnswer2 *string, quickAnswer3 *string, favGame *string, favSubject *string, favBook *string, favQuote *string, favAnimal *string, favAnimation *string, favDevice *string) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditSubject(ctx context.Context, subjectID string, name *model.SubjectEnum, departmentName *model.DeptEnum, weeklyLessons *float64, lessonDuration *float64) (*model.Subject, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditTimeTable(ctx context.Context, ttableID string, day *string, done *bool, lesson1id *string, lesson2id *string, lesson3id *string, lesson4id *string, lesson5id *string, lesson6id *string, lesson7id *string, lesson8id *string, grade *model.GradeEnum) (*model.TimeTable, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditUser(ctx context.Context, userID string, salute *model.SaluteEnum, name *string, iDno *string, nhifNo *string, cell *string, email *string, password *string, poBox *string, poCode *string, district *string, county *string, nationality *string, roles *model.RolesEnum) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditVisitor(ctx context.Context, visitorID string, salute *model.SaluteEnum, name *string, iDno *string, cell *string, email *string, visitingOffice *string, purpose *string, vehicle *string, belongings *string, timeIn *time.Time, timeOut *time.Time, thumbScan *string, visitingFreq *float64, authBy *string) (*model.Visitor, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
