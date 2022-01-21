package querybuilder

import "github.com/goal-web/contracts"

func (this *Builder) SetTX(tx interface{}) contracts.QueryBuilder {
	return this.QueryBuilder.SetTX(tx)
}

func (this *Builder) Create(fields contracts.Fields) interface{} {
	return this.QueryBuilder.Create(fields)
}

func (this *Builder) Insert(values ...contracts.Fields) bool {
	return this.QueryBuilder.Insert(values...)
}

func (this *Builder) Delete() int64 {
	return this.QueryBuilder.Delete()
}

func (this *Builder) Update(fields contracts.Fields) int64 {
	return this.QueryBuilder.Update(fields)
}

func (this *Builder) Get() interface{} {
	return this.QueryBuilder.Get()
}

func (this *Builder) Find(key interface{}) interface{} {
	return this.QueryBuilder.Find(key)
}

func (this *Builder) First() interface{} {
	return this.QueryBuilder.First()
}

func (this *Builder) InsertGetId(values ...contracts.Fields) int64 {
	return this.QueryBuilder.InsertGetId(values...)
}

func (this *Builder) InsertOrIgnore(values ...contracts.Fields) int64 {
	return this.QueryBuilder.InsertOrIgnore(values...)
}

func (this *Builder) InsertOrReplace(values ...contracts.Fields) int64 {
	return this.QueryBuilder.InsertOrReplace(values...)
}

func (this *Builder) FirstOrCreate(values ...contracts.Fields) interface{} {
	return this.QueryBuilder.FirstOrCreate(values...)
}

func (this *Builder) UpdateOrInsert(attributes contracts.Fields, values ...contracts.Fields) bool {
	return this.QueryBuilder.UpdateOrInsert(attributes, values...)
}

func (this *Builder) UpdateOrCreate(attributes contracts.Fields, values ...contracts.Fields) interface{} {
	return this.QueryBuilder.UpdateOrCreate(attributes, values...)
}

func (this *Builder) FirstOrFail() interface{} {
	return this.QueryBuilder.FirstOrFail()
}

func (this *Builder) Count(columns ...string) int64 {
	return this.QueryBuilder.Count(columns...)
}

func (this *Builder) Avg(column string, as ...string) int64 {
	return this.QueryBuilder.Avg(column, as...)
}

func (this *Builder) Sum(column string, as ...string) int64 {
	return this.QueryBuilder.Sum(column, as...)
}

func (this *Builder) Max(column string, as ...string) int64 {
	return this.QueryBuilder.Max(column, as...)
}

func (this *Builder) Min(column string, as ...string) int64 {
	return this.QueryBuilder.Min(column, as...)
}
