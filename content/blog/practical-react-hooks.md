---
title: "React Hooksの実践的な使い方"
date: "2024-06-30"
excerpt: "React Hooksの基本から応用まで、実際のプロジェクトで役立つパターンとアンチパターンを詳しく解説します。"
tags: ["React", "Hooks", "JavaScript"]
---

# React Hooksの実践的な使い方

React Hooksが導入されて以降、関数コンポーネントでの状態管理が格段に改善されました。この記事では、実践的なHooksの使い方をご紹介します。

## 基本的なHooks

### useState

```typescript
import { useState } from 'react'

function Counter() {
  const [count, setCount] = useState(0)
  
  return (
    <div>
      <p>Count: {count}</p>
      <button onClick={() => setCount(count + 1)}>
        Increment
      </button>
    </div>
  )
}
```

### useEffect

```typescript
import { useEffect, useState } from 'react'

function UserProfile({ userId }: { userId: string }) {
  const [user, setUser] = useState(null)
  
  useEffect(() => {
    fetchUser(userId).then(setUser)
  }, [userId])
  
  return <div>{user?.name}</div>
}
```

## カスタムHooks

### useLocalStorage

```typescript
function useLocalStorage<T>(key: string, initialValue: T) {
  const [storedValue, setStoredValue] = useState<T>(() => {
    try {
      const item = window.localStorage.getItem(key)
      return item ? JSON.parse(item) : initialValue
    } catch (error) {
      return initialValue
    }
  })

  const setValue = (value: T | ((val: T) => T)) => {
    try {
      const valueToStore = value instanceof Function ? value(storedValue) : value
      setStoredValue(valueToStore)
      window.localStorage.setItem(key, JSON.stringify(valueToStore))
    } catch (error) {
      console.log(error)
    }
  }

  return [storedValue, setValue] as const
}
```

## ベストプラクティス

1. **依存配列の適切な管理**
2. **メモ化の活用**
3. **カスタムHooksでの関心の分離**

## まとめ

React Hooksを効果的に使用することで、よりクリーンで保守性の高いコードを書くことができます。
