// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go generate in expression/generator; DO NOT EDIT.

package expression

import (
	"github.com/pingcap/tidb/pkg/parser/mysql"
	"github.com/pingcap/tidb/pkg/parser/terror"
	"github.com/pingcap/tidb/pkg/types"
	"github.com/pingcap/tidb/pkg/util/chunk"
)

func (b *builtinAddDatetimeAndDurationSig) vecEvalTime(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	if err := b.args[0].VecEvalTime(ctx, input, result); err != nil {
		return err
	}
	buf0 := result

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalDuration(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf1)

	arg0s := buf0.Times()

	arg1s := buf1.GoDurations()

	resultSlice := result.Times()

	for i := 0; i < n; i++ {

		if result.IsNull(i) {
			continue
		}

		// get arg0 & arg1

		arg0 := arg0s[i]

		arg1 := arg1s[i]

		// calculate

		if arg0.IsZero() {
			result.SetNull(i, true) // fixed: true
			continue
		}

		output, err := arg0.Add(typeCtx(ctx), types.Duration{Duration: arg1, Fsp: -1})

		if err != nil {
			return err
		}

		// commit result

		resultSlice[i] = output

	}
	return nil
}

func (b *builtinAddDatetimeAndDurationSig) vectorized() bool {
	return true
}

func (b *builtinAddDatetimeAndStringSig) vecEvalTime(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	if err := b.args[0].VecEvalTime(ctx, input, result); err != nil {
		return err
	}
	buf0 := result

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalString(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf1)

	arg0s := buf0.Times()

	resultSlice := result.Times()

	for i := 0; i < n; i++ {

		if result.IsNull(i) {
			continue
		}

		// get arg0 & arg1

		arg0 := arg0s[i]

		arg1 := buf1.GetString(i)

		// calculate

		if arg0.IsZero() {
			result.SetNull(i, true) // fixed: true
			continue
		}

		if !isDuration(arg1) {
			result.SetNull(i, true) // fixed: true
			continue
		}
		tc := typeCtx(ctx)
		arg1Duration, _, err := types.ParseDuration(tc, arg1, types.GetFsp(arg1))
		if err != nil {
			if terror.ErrorEqual(err, types.ErrTruncatedWrongVal) {
				tc.AppendWarning(err)
				result.SetNull(i, true) // fixed: true
				continue
			}
			return err
		}

		output, err := arg0.Add(typeCtx(ctx), arg1Duration)

		if err != nil {
			return err
		}

		// commit result

		resultSlice[i] = output

	}
	return nil
}

func (b *builtinAddDatetimeAndStringSig) vectorized() bool {
	return true
}

func (b *builtinAddDurationAndDurationSig) vecEvalDuration(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	if err := b.args[0].VecEvalDuration(ctx, input, result); err != nil {
		return err
	}
	buf0 := result

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalDuration(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf1)

	arg0s := buf0.GoDurations()

	arg1s := buf1.GoDurations()

	resultSlice := result.GoDurations()

	for i := 0; i < n; i++ {

		if result.IsNull(i) {
			continue
		}

		// get arg0 & arg1

		arg0 := arg0s[i]

		arg1 := arg1s[i]

		// calculate

		output, err := types.AddDuration(arg0, arg1)
		if err != nil {
			return err
		}

		// commit result

		resultSlice[i] = output

	}
	return nil
}

func (b *builtinAddDurationAndDurationSig) vectorized() bool {
	return true
}

func (b *builtinAddDurationAndStringSig) vecEvalDuration(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	if err := b.args[0].VecEvalDuration(ctx, input, result); err != nil {
		return err
	}
	buf0 := result

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalString(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf1)

	arg0s := buf0.GoDurations()

	resultSlice := result.GoDurations()

	for i := 0; i < n; i++ {

		if result.IsNull(i) {
			continue
		}

		// get arg0 & arg1

		arg0 := arg0s[i]

		arg1 := buf1.GetString(i)

		// calculate

		if !isDuration(arg1) {
			result.SetNull(i, true) // fixed: true
			continue
		}
		tc := typeCtx(ctx)
		arg1Duration, _, err := types.ParseDuration(tc, arg1, types.GetFsp(arg1))
		if err != nil {
			if terror.ErrorEqual(err, types.ErrTruncatedWrongVal) {
				tc.AppendWarning(err)
				result.SetNull(i, true) // fixed: true
				continue
			}
			return err
		}

		output, err := types.AddDuration(arg0, arg1Duration.Duration)
		if err != nil {
			return err
		}

		// commit result

		resultSlice[i] = output

	}
	return nil
}

func (b *builtinAddDurationAndStringSig) vectorized() bool {
	return true
}

func (b *builtinAddStringAndDurationSig) vecEvalString(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	buf0, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalString(ctx, input, buf0); err != nil {
		return err
	}

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalDuration(ctx, input, buf1); err != nil {
		return err
	}

	result.ReserveString(n)

	arg1s := buf1.GoDurations()

	for i := 0; i < n; i++ {

		if buf0.IsNull(i) || buf1.IsNull(i) {
			result.AppendNull()
			continue
		}

		// get arg0 & arg1

		arg0 := buf0.GetString(i)

		arg1 := arg1s[i]

		// calculate

		tc := typeCtx(ctx)
		fsp1 := b.args[1].GetType(ctx).GetDecimal()
		arg1Duration := types.Duration{Duration: arg1, Fsp: fsp1}
		var output string
		var isNull bool
		if isDuration(arg0) {

			output, err = strDurationAddDuration(tc, arg0, arg1Duration)

			if err != nil {
				if terror.ErrorEqual(err, types.ErrTruncatedWrongVal) {
					tc.AppendWarning(err)
					result.AppendNull() // fixed: false
					continue
				}
				return err
			}
		} else {

			output, isNull, err = strDatetimeAddDuration(tc, arg0, arg1Duration)

			if err != nil {
				return err
			}
			if isNull {
				tc.AppendWarning(err)
				result.AppendNull() // fixed: false
				continue
			}
		}

		// commit result

		result.AppendString(output)

	}
	return nil
}

func (b *builtinAddStringAndDurationSig) vectorized() bool {
	return true
}

func (b *builtinAddStringAndStringSig) vecEvalString(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	buf0, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalString(ctx, input, buf0); err != nil {
		return err
	}

	arg1Type := b.args[1].GetType(ctx)
	if mysql.HasBinaryFlag(arg1Type.GetFlag()) {
		result.ReserveString(n)
		for i := 0; i < n; i++ {
			result.AppendNull()
		}
		return nil
	}

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalString(ctx, input, buf1); err != nil {
		return err
	}

	result.ReserveString(n)

	for i := 0; i < n; i++ {

		if buf0.IsNull(i) || buf1.IsNull(i) {
			result.AppendNull()
			continue
		}

		// get arg0 & arg1

		arg0 := buf0.GetString(i)

		arg1 := buf1.GetString(i)

		// calculate

		tc := typeCtx(ctx)
		arg1Duration, _, err := types.ParseDuration(tc, arg1, getFsp4TimeAddSub(arg1))
		if err != nil {
			if terror.ErrorEqual(err, types.ErrTruncatedWrongVal) {
				tc.AppendWarning(err)
				result.AppendNull() // fixed: false
				continue
			}
			return err
		}

		var output string
		var isNull bool
		if isDuration(arg0) {

			output, err = strDurationAddDuration(tc, arg0, arg1Duration)

			if err != nil {
				if terror.ErrorEqual(err, types.ErrTruncatedWrongVal) {
					tc.AppendWarning(err)
					result.AppendNull() // fixed: false
					continue
				}
				return err
			}
		} else {

			output, isNull, err = strDatetimeAddDuration(tc, arg0, arg1Duration)

			if err != nil {
				return err
			}
			if isNull {
				tc.AppendWarning(err)
				result.AppendNull() // fixed: false
				continue
			}
		}

		// commit result

		result.AppendString(output)

	}
	return nil
}

func (b *builtinAddStringAndStringSig) vectorized() bool {
	return true
}

func (b *builtinAddDateAndDurationSig) vecEvalString(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	buf0, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalTime(ctx, input, buf0); err != nil {
		return err
	}

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalDuration(ctx, input, buf1); err != nil {
		return err
	}

	result.ReserveString(n)

	arg0s := buf0.Times()

	arg1s := buf1.GoDurations()

	for i := 0; i < n; i++ {

		if buf0.IsNull(i) || buf1.IsNull(i) {
			result.AppendNull()
			continue
		}

		// get arg0 & arg1

		arg0 := arg0s[i]

		arg1 := arg1s[i]

		// calculate

		if arg0.IsZero() {
			result.AppendNull() // fixed: false
			continue
		}

		fsp1 := b.args[1].GetType(ctx).GetDecimal()
		arg1Duration := types.Duration{Duration: arg1, Fsp: fsp1}
		tc := typeCtx(ctx)
		arg0.SetType(mysql.TypeDatetime)

		res, err := arg0.Add(tc, arg1Duration)

		if err != nil {
			tc.AppendWarning(err)
			result.AppendNull() // fixed: false
			continue
		}

		output := res.String()

		// commit result

		result.AppendString(output)

	}
	return nil
}

func (b *builtinAddDateAndDurationSig) vectorized() bool {
	return true
}

func (b *builtinAddDateAndStringSig) vecEvalString(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	buf0, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalTime(ctx, input, buf0); err != nil {
		return err
	}

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalString(ctx, input, buf1); err != nil {
		return err
	}

	result.ReserveString(n)

	arg0s := buf0.Times()

	for i := 0; i < n; i++ {

		if buf0.IsNull(i) || buf1.IsNull(i) {
			result.AppendNull()
			continue
		}

		// get arg0 & arg1

		arg0 := arg0s[i]

		arg1 := buf1.GetString(i)

		// calculate

		if arg0.IsZero() {
			result.AppendNull() // fixed: false
			continue
		}

		if !isDuration(arg1) {
			result.AppendNull() // fixed: false
			continue
		}
		tc := typeCtx(ctx)
		arg1Duration, _, err := types.ParseDuration(tc, arg1, getFsp4TimeAddSub(arg1))
		if err != nil {
			if terror.ErrorEqual(err, types.ErrTruncatedWrongVal) {
				tc.AppendWarning(err)
				result.AppendNull() // fixed: false
				continue
			}
			return err
		}

		arg0.SetType(mysql.TypeDatetime)

		res, err := arg0.Add(tc, arg1Duration)

		if err != nil {
			tc.AppendWarning(err)
			result.AppendNull() // fixed: false
			continue
		}

		output := res.String()

		// commit result

		result.AppendString(output)

	}
	return nil
}

func (b *builtinAddDateAndStringSig) vectorized() bool {
	return true
}

func (b *builtinAddTimeDateTimeNullSig) vecEvalTime(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	result.ResizeTime(n, true)

	return nil
}

func (b *builtinAddTimeDateTimeNullSig) vectorized() bool {
	return true
}

func (b *builtinAddTimeStringNullSig) vecEvalString(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	result.ReserveString(n)
	for i := 0; i < n; i++ {
		result.AppendNull()
	}

	return nil
}

func (b *builtinAddTimeStringNullSig) vectorized() bool {
	return true
}

func (b *builtinAddTimeDurationNullSig) vecEvalDuration(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	result.ResizeGoDuration(n, true)

	return nil
}

func (b *builtinAddTimeDurationNullSig) vectorized() bool {
	return true
}

func (b *builtinSubDatetimeAndDurationSig) vecEvalTime(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	if err := b.args[0].VecEvalTime(ctx, input, result); err != nil {
		return err
	}
	buf0 := result

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalDuration(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf1)

	arg0s := buf0.Times()

	arg1s := buf1.GoDurations()

	resultSlice := result.Times()

	for i := 0; i < n; i++ {

		if result.IsNull(i) {
			continue
		}

		// get arg0 & arg1

		arg0 := arg0s[i]

		arg1 := arg1s[i]

		// calculate

		if arg0.IsZero() {
			result.SetNull(i, true) // fixed: true
			continue
		}

		tc := typeCtx(ctx)
		arg1Duration := types.Duration{Duration: arg1, Fsp: -1}
		output, err := arg0.Add(tc, arg1Duration.Neg())

		if err != nil {
			return err
		}

		// commit result

		resultSlice[i] = output

	}
	return nil
}

func (b *builtinSubDatetimeAndDurationSig) vectorized() bool {
	return true
}

func (b *builtinSubDatetimeAndStringSig) vecEvalTime(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	if err := b.args[0].VecEvalTime(ctx, input, result); err != nil {
		return err
	}
	buf0 := result

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalString(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf1)

	arg0s := buf0.Times()

	resultSlice := result.Times()

	for i := 0; i < n; i++ {

		if result.IsNull(i) {
			continue
		}

		// get arg0 & arg1

		arg0 := arg0s[i]

		arg1 := buf1.GetString(i)

		// calculate

		if arg0.IsZero() {
			result.SetNull(i, true) // fixed: true
			continue
		}

		if !isDuration(arg1) {
			result.SetNull(i, true) // fixed: true
			continue
		}
		tc := typeCtx(ctx)
		arg1Duration, _, err := types.ParseDuration(tc, arg1, types.GetFsp(arg1))
		if err != nil {
			if terror.ErrorEqual(err, types.ErrTruncatedWrongVal) {
				tc.AppendWarning(err)
				result.SetNull(i, true) // fixed: true
				continue
			}
			return err
		}
		output, err := arg0.Add(typeCtx(ctx), arg1Duration.Neg())

		if err != nil {
			return err
		}

		// commit result

		resultSlice[i] = output

	}
	return nil
}

func (b *builtinSubDatetimeAndStringSig) vectorized() bool {
	return true
}

func (b *builtinSubDurationAndDurationSig) vecEvalDuration(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	if err := b.args[0].VecEvalDuration(ctx, input, result); err != nil {
		return err
	}
	buf0 := result

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalDuration(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf1)

	arg0s := buf0.GoDurations()

	arg1s := buf1.GoDurations()

	resultSlice := result.GoDurations()

	for i := 0; i < n; i++ {

		if result.IsNull(i) {
			continue
		}

		// get arg0 & arg1

		arg0 := arg0s[i]

		arg1 := arg1s[i]

		// calculate

		output, err := types.SubDuration(arg0, arg1)
		if err != nil {
			return err
		}

		// commit result

		resultSlice[i] = output

	}
	return nil
}

func (b *builtinSubDurationAndDurationSig) vectorized() bool {
	return true
}

func (b *builtinSubDurationAndStringSig) vecEvalDuration(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	if err := b.args[0].VecEvalDuration(ctx, input, result); err != nil {
		return err
	}
	buf0 := result

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalString(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf1)

	arg0s := buf0.GoDurations()

	resultSlice := result.GoDurations()

	for i := 0; i < n; i++ {

		if result.IsNull(i) {
			continue
		}

		// get arg0 & arg1

		arg0 := arg0s[i]

		arg1 := buf1.GetString(i)

		// calculate

		if !isDuration(arg1) {
			result.SetNull(i, true) // fixed: true
			continue
		}
		tc := typeCtx(ctx)
		arg1Duration, _, err := types.ParseDuration(tc, arg1, types.GetFsp(arg1))
		if err != nil {
			if terror.ErrorEqual(err, types.ErrTruncatedWrongVal) {
				tc.AppendWarning(err)
				result.SetNull(i, true) // fixed: true
				continue
			}
			return err
		}

		output, err := types.SubDuration(arg0, arg1Duration.Duration)
		if err != nil {
			return err
		}

		// commit result

		resultSlice[i] = output

	}
	return nil
}

func (b *builtinSubDurationAndStringSig) vectorized() bool {
	return true
}

func (b *builtinSubStringAndDurationSig) vecEvalString(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	buf0, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalString(ctx, input, buf0); err != nil {
		return err
	}

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalDuration(ctx, input, buf1); err != nil {
		return err
	}

	result.ReserveString(n)

	arg1s := buf1.GoDurations()

	for i := 0; i < n; i++ {

		if buf0.IsNull(i) || buf1.IsNull(i) {
			result.AppendNull()
			continue
		}

		// get arg0 & arg1

		arg0 := buf0.GetString(i)

		arg1 := arg1s[i]

		// calculate

		tc := typeCtx(ctx)
		fsp1 := b.args[1].GetType(ctx).GetDecimal()
		arg1Duration := types.Duration{Duration: arg1, Fsp: fsp1}
		var output string
		var isNull bool
		if isDuration(arg0) {

			output, err = strDurationSubDuration(tc, arg0, arg1Duration)

			if err != nil {
				if terror.ErrorEqual(err, types.ErrTruncatedWrongVal) {
					tc.AppendWarning(err)
					result.AppendNull() // fixed: false
					continue
				}
				return err
			}
		} else {

			output, isNull, err = strDatetimeSubDuration(tc, arg0, arg1Duration)

			if err != nil {
				return err
			}
			if isNull {
				tc.AppendWarning(err)
				result.AppendNull() // fixed: false
				continue
			}
		}

		// commit result

		result.AppendString(output)

	}
	return nil
}

func (b *builtinSubStringAndDurationSig) vectorized() bool {
	return true
}

func (b *builtinSubStringAndStringSig) vecEvalString(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	buf0, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalString(ctx, input, buf0); err != nil {
		return err
	}

	arg1Type := b.args[1].GetType(ctx)
	if mysql.HasBinaryFlag(arg1Type.GetFlag()) {
		result.ReserveString(n)
		for i := 0; i < n; i++ {
			result.AppendNull()
		}
		return nil
	}

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalString(ctx, input, buf1); err != nil {
		return err
	}

	result.ReserveString(n)

	for i := 0; i < n; i++ {

		if buf0.IsNull(i) || buf1.IsNull(i) {
			result.AppendNull()
			continue
		}

		// get arg0 & arg1

		arg0 := buf0.GetString(i)

		arg1 := buf1.GetString(i)

		// calculate

		tc := typeCtx(ctx)
		arg1Duration, _, err := types.ParseDuration(tc, arg1, getFsp4TimeAddSub(arg1))
		if err != nil {
			if terror.ErrorEqual(err, types.ErrTruncatedWrongVal) {
				tc.AppendWarning(err)
				result.AppendNull() // fixed: false
				continue
			}
			return err
		}

		var output string
		var isNull bool
		if isDuration(arg0) {

			output, err = strDurationSubDuration(tc, arg0, arg1Duration)

			if err != nil {
				if terror.ErrorEqual(err, types.ErrTruncatedWrongVal) {
					tc.AppendWarning(err)
					result.AppendNull() // fixed: false
					continue
				}
				return err
			}
		} else {

			output, isNull, err = strDatetimeSubDuration(tc, arg0, arg1Duration)

			if err != nil {
				return err
			}
			if isNull {
				tc.AppendWarning(err)
				result.AppendNull() // fixed: false
				continue
			}
		}

		// commit result

		result.AppendString(output)

	}
	return nil
}

func (b *builtinSubStringAndStringSig) vectorized() bool {
	return true
}

func (b *builtinSubDateAndDurationSig) vecEvalString(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	buf0, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalTime(ctx, input, buf0); err != nil {
		return err
	}

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalDuration(ctx, input, buf1); err != nil {
		return err
	}

	result.ReserveString(n)

	arg0s := buf0.Times()

	arg1s := buf1.GoDurations()

	for i := 0; i < n; i++ {

		if buf0.IsNull(i) || buf1.IsNull(i) {
			result.AppendNull()
			continue
		}

		// get arg0 & arg1

		arg0 := arg0s[i]

		arg1 := arg1s[i]

		// calculate

		if arg0.IsZero() {
			result.AppendNull() // fixed: false
			continue
		}

		fsp1 := b.args[1].GetType(ctx).GetDecimal()
		arg1Duration := types.Duration{Duration: arg1, Fsp: fsp1}
		tc := typeCtx(ctx)
		arg0.SetType(mysql.TypeDatetime)

		res, err := arg0.Add(tc, arg1Duration.Neg())

		if err != nil {
			tc.AppendWarning(err)
			result.AppendNull() // fixed: false
			continue
		}

		output := res.String()

		// commit result

		result.AppendString(output)

	}
	return nil
}

func (b *builtinSubDateAndDurationSig) vectorized() bool {
	return true
}

func (b *builtinSubDateAndStringSig) vecEvalString(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	buf0, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalTime(ctx, input, buf0); err != nil {
		return err
	}

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalString(ctx, input, buf1); err != nil {
		return err
	}

	result.ReserveString(n)

	arg0s := buf0.Times()

	for i := 0; i < n; i++ {

		if buf0.IsNull(i) || buf1.IsNull(i) {
			result.AppendNull()
			continue
		}

		// get arg0 & arg1

		arg0 := arg0s[i]

		arg1 := buf1.GetString(i)

		// calculate

		if arg0.IsZero() {
			result.AppendNull() // fixed: false
			continue
		}

		if !isDuration(arg1) {
			result.AppendNull() // fixed: false
			continue
		}
		tc := typeCtx(ctx)
		arg1Duration, _, err := types.ParseDuration(tc, arg1, getFsp4TimeAddSub(arg1))
		if err != nil {
			if terror.ErrorEqual(err, types.ErrTruncatedWrongVal) {
				tc.AppendWarning(err)
				result.AppendNull() // fixed: false
				continue
			}
			return err
		}

		arg0.SetType(mysql.TypeDatetime)

		res, err := arg0.Add(tc, arg1Duration.Neg())

		if err != nil {
			tc.AppendWarning(err)
			result.AppendNull() // fixed: false
			continue
		}

		output := res.String()

		// commit result

		result.AppendString(output)

	}
	return nil
}

func (b *builtinSubDateAndStringSig) vectorized() bool {
	return true
}

func (b *builtinSubTimeDateTimeNullSig) vecEvalTime(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	result.ResizeTime(n, true)

	return nil
}

func (b *builtinSubTimeDateTimeNullSig) vectorized() bool {
	return true
}

func (b *builtinSubTimeStringNullSig) vecEvalString(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	result.ReserveString(n)
	for i := 0; i < n; i++ {
		result.AppendNull()
	}

	return nil
}

func (b *builtinSubTimeStringNullSig) vectorized() bool {
	return true
}

func (b *builtinSubTimeDurationNullSig) vecEvalDuration(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	result.ResizeGoDuration(n, true)

	return nil
}

func (b *builtinSubTimeDurationNullSig) vectorized() bool {
	return true
}

func (b *builtinNullTimeDiffSig) vecEvalDuration(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	result.ResizeGoDuration(n, true)
	return nil
}

func (b *builtinNullTimeDiffSig) vectorized() bool {
	return true
}

func (b *builtinTimeStringTimeDiffSig) vecEvalDuration(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	result.ResizeGoDuration(n, false)
	r64s := result.GoDurations()
	buf0, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)

	if err := b.args[0].VecEvalTime(ctx, input, buf0); err != nil {
		return err
	}
	if err := b.args[1].VecEvalString(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf0, buf1)
	arg0 := buf0.Times()
	tc := typeCtx(ctx)
	for i := 0; i < n; i++ {
		if result.IsNull(i) {
			continue
		}
		lhsTime := arg0[i]
		_, rhsTime, rhsIsDuration, err := convertStringToDuration(tc, buf1.GetString(i), b.tp.GetDecimal())
		if err != nil {
			return err
		}
		if rhsIsDuration {
			result.SetNull(i, true)
			continue
		}
		d, isNull, err := calculateTimeDiff(tc, lhsTime, rhsTime)
		if err != nil {
			return err
		}
		if isNull {
			result.SetNull(i, true)
			continue
		}
		r64s[i] = d.Duration
	}
	return nil
}

func (b *builtinTimeStringTimeDiffSig) vectorized() bool {
	return true
}

func (b *builtinDurationStringTimeDiffSig) vecEvalDuration(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	result.ResizeGoDuration(n, false)
	r64s := result.GoDurations()
	buf0 := result
	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)

	if err := b.args[0].VecEvalDuration(ctx, input, buf0); err != nil {
		return err
	}
	if err := b.args[1].VecEvalString(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf1)
	arg0 := buf0.GoDurations()
	var (
		lhs types.Duration
		rhs types.Duration
	)
	tc := typeCtx(ctx)
	for i := 0; i < n; i++ {
		if result.IsNull(i) {
			continue
		}
		lhs.Duration = arg0[i]
		rhsDur, _, rhsIsDuration, err := convertStringToDuration(tc, buf1.GetString(i), b.tp.GetDecimal())
		if err != nil {
			return err
		}
		if !rhsIsDuration {
			result.SetNull(i, true)
			continue
		}
		rhs = rhsDur
		d, isNull, err := calculateDurationTimeDiff(ctx, lhs, rhs)
		if err != nil {
			return err
		}
		if isNull {
			result.SetNull(i, true)
			continue
		}
		r64s[i] = d.Duration
	}
	return nil
}

func (b *builtinDurationStringTimeDiffSig) vectorized() bool {
	return true
}

func (b *builtinDurationDurationTimeDiffSig) vecEvalDuration(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	result.ResizeGoDuration(n, false)
	r64s := result.GoDurations()
	buf0 := result
	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)

	if err := b.args[0].VecEvalDuration(ctx, input, buf0); err != nil {
		return err
	}
	if err := b.args[1].VecEvalDuration(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf1)
	arg0 := buf0.GoDurations()
	arg1 := buf1.GoDurations()
	var (
		lhs types.Duration
		rhs types.Duration
	)
	for i := 0; i < n; i++ {
		if result.IsNull(i) {
			continue
		}
		lhs.Duration = arg0[i]
		rhs.Duration = arg1[i]
		d, isNull, err := calculateDurationTimeDiff(ctx, lhs, rhs)
		if err != nil {
			return err
		}
		if isNull {
			result.SetNull(i, true)
			continue
		}
		r64s[i] = d.Duration
	}
	return nil
}

func (b *builtinDurationDurationTimeDiffSig) vectorized() bool {
	return true
}

func (b *builtinStringTimeTimeDiffSig) vecEvalDuration(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	result.ResizeGoDuration(n, false)
	r64s := result.GoDurations()
	buf0, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)

	if err := b.args[0].VecEvalString(ctx, input, buf0); err != nil {
		return err
	}
	if err := b.args[1].VecEvalTime(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf0, buf1)
	arg1 := buf1.Times()
	tc := typeCtx(ctx)
	for i := 0; i < n; i++ {
		if result.IsNull(i) {
			continue
		}
		_, lhsTime, lhsIsDuration, err := convertStringToDuration(tc, buf0.GetString(i), b.tp.GetDecimal())
		if err != nil {
			return err
		}
		if lhsIsDuration {
			result.SetNull(i, true)
			continue
		}
		rhsTime := arg1[i]
		d, isNull, err := calculateTimeDiff(tc, lhsTime, rhsTime)
		if err != nil {
			return err
		}
		if isNull {
			result.SetNull(i, true)
			continue
		}
		r64s[i] = d.Duration
	}
	return nil
}

func (b *builtinStringTimeTimeDiffSig) vectorized() bool {
	return true
}

func (b *builtinStringDurationTimeDiffSig) vecEvalDuration(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	result.ResizeGoDuration(n, false)
	r64s := result.GoDurations()
	buf1 := result
	buf0, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)

	if err := b.args[0].VecEvalString(ctx, input, buf0); err != nil {
		return err
	}
	if err := b.args[1].VecEvalDuration(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf0)
	arg1 := buf1.GoDurations()
	var (
		lhs types.Duration
		rhs types.Duration
	)
	tc := typeCtx(ctx)
	for i := 0; i < n; i++ {
		if result.IsNull(i) {
			continue
		}
		lhsDur, _, lhsIsDuration, err := convertStringToDuration(tc, buf0.GetString(i), b.tp.GetDecimal())
		if err != nil {
			return err
		}
		if !lhsIsDuration {
			result.SetNull(i, true)
			continue
		}
		lhs = lhsDur
		rhs.Duration = arg1[i]
		d, isNull, err := calculateDurationTimeDiff(ctx, lhs, rhs)
		if err != nil {
			return err
		}
		if isNull {
			result.SetNull(i, true)
			continue
		}
		r64s[i] = d.Duration
	}
	return nil
}

func (b *builtinStringDurationTimeDiffSig) vectorized() bool {
	return true
}

func (b *builtinStringStringTimeDiffSig) vecEvalDuration(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	result.ResizeGoDuration(n, false)
	r64s := result.GoDurations()
	buf0, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)

	if err := b.args[0].VecEvalString(ctx, input, buf0); err != nil {
		return err
	}
	if err := b.args[1].VecEvalString(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf0, buf1)
	tc := typeCtx(ctx)
	for i := 0; i < n; i++ {
		if result.IsNull(i) {
			continue
		}
		lhsDur, lhsTime, lhsIsDuration, err := convertStringToDuration(tc, buf0.GetString(i), b.tp.GetDecimal())
		if err != nil {
			return err
		}
		rhsDur, rhsTime, rhsIsDuration, err := convertStringToDuration(tc, buf1.GetString(i), b.tp.GetDecimal())
		if err != nil {
			return err
		}
		if lhsIsDuration != rhsIsDuration {
			result.SetNull(i, true)
			continue
		}
		var (
			d      types.Duration
			isNull bool
		)
		if lhsIsDuration {
			d, isNull, err = calculateDurationTimeDiff(ctx, lhsDur, rhsDur)
		} else {
			d, isNull, err = calculateTimeDiff(tc, lhsTime, rhsTime)
		}
		if err != nil {
			return err
		}
		if isNull {
			result.SetNull(i, true)
			continue
		}
		r64s[i] = d.Duration
	}
	return nil
}

func (b *builtinStringStringTimeDiffSig) vectorized() bool {
	return true
}

func (b *builtinTimeTimeTimeDiffSig) vecEvalDuration(ctx EvalContext, input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	result.ResizeGoDuration(n, false)
	r64s := result.GoDurations()
	buf0, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)

	buf1, err := b.bufAllocator.get()
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)

	if err := b.args[0].VecEvalTime(ctx, input, buf0); err != nil {
		return err
	}
	if err := b.args[1].VecEvalTime(ctx, input, buf1); err != nil {
		return err
	}

	result.MergeNulls(buf0, buf1)
	arg0 := buf0.Times()
	arg1 := buf1.Times()
	tc := typeCtx(ctx)
	for i := 0; i < n; i++ {
		if result.IsNull(i) {
			continue
		}
		lhsTime := arg0[i]
		rhsTime := arg1[i]
		d, isNull, err := calculateTimeDiff(tc, lhsTime, rhsTime)
		if err != nil {
			return err
		}
		if isNull {
			result.SetNull(i, true)
			continue
		}
		r64s[i] = d.Duration
	}
	return nil
}

func (b *builtinTimeTimeTimeDiffSig) vectorized() bool {
	return true
}