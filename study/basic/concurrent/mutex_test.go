package concurrent

import "testing"

func TestMutex1(t *testing.T) {
	mutex.Lock()
	defer mutex.Unlock()
	// 你的代码
}

func TestRwMutex(t *testing.T) {
	// 加读锁
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	// 也可以加写锁
	rwMutex.Lock()
	defer rwMutex.Unlock()
}

// 不可重入例子
func TestFailed1(t *testing.T) {
	mutex.Lock()
	defer mutex.Unlock()
	mutex.Lock()
	defer mutex.Unlock()
}

// 不可升级
func TestFailed2(t *testing.T) {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	mutex.Lock()
	defer mutex.Unlock()
}

//func TestFailed3() {
//
//}
